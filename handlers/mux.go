package handlers

import "github.com/gorilla/mux"

// NewMux ...
func NewMux() *mux.Router {
	// Create router
	mux := mux.NewRouter()

	// Register route handlers
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/echo/{word}", Echo)

	return mux
}
