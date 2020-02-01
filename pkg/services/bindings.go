// +build ziplinegen

package services

import (
	zipline "github.com/bilal-bhatti/zipline/pkg"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	mux := chi.NewRouter()

	// bind contacts
	mux.Post("/contacts", zipline.Post(InitContactsService().Create))
	mux.Get("/contacts/{id}", zipline.Get(InitContactsService().GetOne))

	// bind things
	mux.Post("/things", zipline.Post(InitThingsService().Create))
	mux.Get("/things/{id}", zipline.Get(InitThingsService().GetOne))

	mux.Post("/echo", zipline.Post(Echo))

	return mux
}
