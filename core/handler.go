package core

import (
	"net/http"
	"strings"
)

type Handler struct {
	APIHandler  http.Handler
	FileHandler http.Handler
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/api"):
		http.StripPrefix("/api", h.APIHandler).ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/"):
		h.FileHandler.ServeHTTP(w, r)
	}
}
