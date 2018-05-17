package core

import (
	"github.com/gorilla/mux"
	"github.com/icheckteam/explorer/store"
)

// API ...
type API struct {
	store *store.Store
}

func NewAPI(s *store.Store) *API {
	return &API{store: s}
}

func (a *API) RegisterRoutes(mux *mux.Router) {

	// Blocks ....
	mux.HandleFunc("/blocks/hash/{hash}", a.handleGetBlockHash)
	mux.HandleFunc("/blocks/height/{height}", a.handleGetBlockHeight)
	mux.HandleFunc("/blocks", a.handleGetBlocks)

	// Txns
	mux.HandleFunc("/txs/{hash}", a.handleGetTxnHash)
	mux.HandleFunc("/txs", a.handleGetTxs)

	// Assets
	mux.HandleFunc("/assets", a.handleGetAssets)
	mux.HandleFunc("/assets/{id}", a.handleGetAsset)
	mux.HandleFunc("/assets/{id}/accounts", a.handleGetAssetAccounts)
	mux.HandleFunc("/assets/{id}/trasfers", a.handleGetAssetTransfers)
	mux.HandleFunc("/assets/{id}/txs", a.handleGetAssetTxs)
	mux.HandleFunc("/assets/{id}/issues", a.handleGetAssetIssues)
	mux.HandleFunc("/assets/{id}/history", a.handleGetAssetHistory)
}
