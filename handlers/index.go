package handlers

import (
	"io"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	io.WriteString(w, "Try /echo/{word}")
}
