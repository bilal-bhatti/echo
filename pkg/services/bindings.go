// +build ziplinegen

package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	mux := chi.NewRouter()

	// bind contacts
	mux.Post("/contacts", zipline.Post(ContactsService.Create))
	mux.Get("/contacts/{id}", zipline.Get(ContactsService.GetOne))

	// bind things
	mux.Post("/things", zipline.Post(ThingsService.Create))
	mux.Get("/things/{id}", zipline.Get(ThingsService.GetOne))

	mux.Post("/echo", zipline.Post(Echo))

	return mux
}

var zipline ZiplineTemplate

type ZiplineTemplate struct {
	ReturnResponseAndError func() (interface{}, error)
}

func (z ZiplineTemplate) Post(i interface{}) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		response, err := z.ReturnResponseAndError()
		if err != nil {
			// write error response
			// internal error
			panic(err)
		}

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			// write error response
			panic(err)
		}
	}
}

func (z ZiplineTemplate) Get(i interface{}) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		response, err := z.ReturnResponseAndError()
		if err != nil {
			// write error response
			// internal error
			panic(err)
		}

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			// write error response
			panic(err)
		}
	}
}
