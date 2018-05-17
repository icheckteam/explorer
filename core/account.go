package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetTxs ...
func (s *API) handleGetAccountAssets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit, skip, err := getLimitSkip(r)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	b, err := s.store.GetAccountAssets(vars["id"], limit, skip)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}
