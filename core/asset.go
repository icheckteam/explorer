package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetTxs ...
func (s *API) handleGetAssets(w http.ResponseWriter, r *http.Request) {
	limit, skip, err := getLimitSkip(r)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	b, err := s.store.GetAssets(limit, skip)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}

// handleGetTxs ...
func (s *API) handleGetAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := s.store.GetAsset(vars["id"])
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}
