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
	mux.HandleFunc("/assets", a.handleGetAssets)
	mux.HandleFunc("/assets/{id}", a.handleGetAsset)
	mux.HandleFunc("/assets/{id}/history/{name}", a.handleGetAssetHistory)
}
