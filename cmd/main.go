package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/icheckteam/explorer/core"
	"github.com/icheckteam/explorer/sdk"
	"github.com/icheckteam/explorer/store"
	"github.com/icheckteam/explorer/types"
	"github.com/icheckteam/ichain/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmserver "github.com/tendermint/tendermint/rpc/lib/server"
	"github.com/tendermint/tmlibs/cli"
	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
	mgo "gopkg.in/mgo.v2"
)

const (
	flagListenAddr = "laddr"
	flagCORS       = "cors"
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

	rootCmd.AddCommand(GetStartRestServerCmd(cdc, s))

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "E", os.ExpandEnv("$HOME/.explorerli"))
	executor.Execute()
}

func GetStartCmd(cdc *wire.Codec, s *store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "start-indexer",
		Short: "start-indexer",
		Run: func(cmd *cobra.Command, args []string) {
			currentHeight, _ := s.GetCurrentHeight()
			currentHeight += 1
			for {
				fmt.Printf("block height: %d", currentHeight)
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

func GetStartRestServerCmd(cdc *wire.Codec, s *store.Store) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rest-server",
		Short: "rest-server",
		RunE:  startRESTServerFn(s),
	}
	cmd.Flags().StringP(flagListenAddr, "a", "tcp://localhost:1318", "Address for server to listen on")
	cmd.Flags().String(flagCORS, "", "Set to domains that can make CORS requests (* for all)")
	cmd.Flags().StringP(client.FlagChainID, "c", "", "ID of chain we connect to")
	cmd.Flags().StringP(client.FlagNode, "n", "tcp://localhost:46657", "Node to connect to")
	return cmd
}

func startRESTServerFn(s *store.Store) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		listenAddr := viper.GetString(flagListenAddr)
		handler := createHandler(s)
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).
			With("module", "rest-server")
		listener, err := tmserver.StartHTTPServer(listenAddr, handler, logger)
		if err != nil {
			return err
		}

		// Wait forever and cleanup
		cmn.TrapSignal(func() {
			err := listener.Close()
			logger.Error("Error closing listener", "err", err)
		})
		return nil
	}
}

func createHandler(s *store.Store) http.Handler {
	r := mux.NewRouter()
	a := core.NewAPI(s)
	a.RegisterRoutes(r)
	return core.Handler{
		APIHandler:  r,
		FileHandler: core.NewFileHandler("./dashboard/build"),
	}
}
