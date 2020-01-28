package services

import (
	zipline "github.com/bilal-bhatti/zipline/pkg"

	"github.com/go-chi/chi"
)

// NewRouter ....
func NewRouter() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/contacts", zipline.Post(InitContactsService().Create))
	mux.Get("/contacts/{id}", zipline.Get(InitContactsService().GetOne))

	// bind things
	things := InitThingsService()
	mux.Post("/things", zipline.Post(things.Create))
	mux.Get("/things/{id}", zipline.Get(things.GetOne))

	mux.Post("/echo", zipline.Post(Echo))

	return mux
}
