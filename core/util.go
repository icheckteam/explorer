package core

import (
	"encoding/json"
	"net/http"

	"github.com/icheckteam/explorer/types"
)

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteErrorJSON(w http.ResponseWriter, err error) {
	v := types.ErrorResponse{
		Error: types.Error{Msg: err.Error()},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return
}
