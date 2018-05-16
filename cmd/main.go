package main

import (
	"log"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/icheckteam/explorer/sdk"
	"github.com/icheckteam/explorer/store"
	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/app"
	"github.com/spf13/cobra"
	"github.com/tendermint/tmlibs/cli"
	mgo "gopkg.in/mgo.v2"
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
				log.Printf("block height: %d", currentHeight)
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

				if res.Block.NumTxs > 0 {
					for i, tx := range res.Block.Data.Txs {
						ptx, err := parseTx(cdc, tx)
						if err != nil {
							panic(err)
						}
						txInfo := types.TxInfo{
							Hash:   tx.String(),
							Time:   res.Block.Time,
							Height: res.Block.Height,
							Tx:     ptx,
							Index:  int64(i),
						}

						s.InsertTxn(txInfo)
					}
				}

				currentHeight++

			}
		},
	}
}

func parseTx(cdc *wire.Codec, txBytes []byte) (csdk.Tx, error) {
	var tx csdk.StdTx
	err := cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
