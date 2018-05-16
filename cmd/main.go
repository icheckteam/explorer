package main

import (
	"log"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/icheckteam/explorer/sdk"
	"github.com/icheckteam/explorer/store"
	"github.com/icheckteam/ichain/app"
	"github.com/spf13/cobra"
	mgo "gopkg.in/mgo.v2"

	"github.com/tendermint/tmlibs/cli"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
	}
)

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	s := store.NewStore(session.DB("ichain"))

	// add proxy, version and key info
	rootCmd.AddCommand(
		client.GetCommands(
			GetStartCmd(cdc, s),
		)...)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "E", os.ExpandEnv("$HOME/.explorerli"))
	executor.Execute()
}

func GetStartCmd(cdc *wire.Codec, s *store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start",
		Run: func(cmd *cobra.Command, args []string) {
			currentHeight, _ := s.GetCurrentHeight()
			currentHeight += 1
			for {
				ichainSdk := sdk.NewIchainSdk(context.NewCoreContextFromViper())
				res, err := ichainSdk.GetBlock(currentHeight)
				if err != nil {
					time.Sleep(time.Second * 1)
					continue
				}

				// insert block
				s.InsertBlock(res.Block)

				// update next block hash
				if res.Block.Height > 1 {
					s.UpdateNextBlockHash(res.Block.Hash().String(), res.Block.Height-1)
				}

				log.Printf("run %d", res.Block.Header.Height)
				currentHeight++

			}
		},
	}
}
