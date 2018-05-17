package core

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handleGetBlockHash ...
func (s *API) handleGetBlockHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := s.store.GetBlockHash(vars["hash"])
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}

// handleGetBlocks ...
func (s *API) handleGetBlockHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height, err := strconv.Atoi(vars["height"])
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	b, err := s.store.GetBlockHeight(int64(height))
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}

// handleGetBlocks ...
func (s *API) handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	limit, skip, err := getLimitSkip(r)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	b, err := s.store.GetBlocks(limit, skip)
	if err != nil {
		WriteErrorJSON(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, b)
}

func getLimitSkip(r *http.Request) (int, int, error) {
	var err error
	var limit, skip int
	vars := r.URL.Query()
	if vars.Get("limit") != "" {
		limit, err = strconv.Atoi(vars.Get("limit"))
		if err != nil {
			return 0, 0, err
		}
	} else {
		limit = 10
	}

	if limit >= 200 {
		return 0, 0, errors.New("max limit size")
	}

	if vars.Get("skip") != "" {
		skip, err = strconv.Atoi(vars.Get("skip"))
		if err != nil {
			return 0, 0, err
		}
	} else {
		skip = 0
	}
	return limit, skip, nil
}
