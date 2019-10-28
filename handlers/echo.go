package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Echo ...
func Echo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Replace i with o in the word
	io.WriteString(w, strings.Replace(vars["word"], "i", "o", -1))
}
