//go:build ziplinegen
// +build ziplinegen

package services

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewRouter(mux *chi.Mux) *chi.Mux {
	// bind contacts
	mux.Post("/api/contacts", z.Post(ContactsService.Create, z.Resolve, z.Body))
	mux.Get("/api/contacts/{id}", z.Get(ContactsService.GetOne, z.Path))

	// // bind things
	mux.Post("/api/things", z.Post(ThingsService.Create, z.Body))
	mux.Get("/api/things/{id}", z.Get(ThingsService.GetOne, z.Path))

	// mux.Post("/echo", z.Post(Echo))

	return mux
}

var z ZiplineTemplate

type ZiplineTemplate struct {
	ReturnResponseAndError func() (interface{}, error)
	ReturnError            func() error
	Resolve                func() (ZiplineTemplate, error)
	DevNull                func(i ...interface{})
}

func (z ZiplineTemplate) Body(w http.ResponseWriter, r *http.Request) {
	var err error
	defer io.Copy(ioutil.Discard, r.Body)

	name := ZiplineTemplate{}
	err = json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (z ZiplineTemplate) Path(kind string, w http.ResponseWriter, r *http.Request) {
	switch kind {
	case "string":
		name := chi.URLParam(r, "name")
		z.DevNull(name)
	case "int":
		name, err := strconv.Atoi(chi.URLParam(r, "name"))
		if err != nil {
			panic(err)
		}
		z.DevNull(name)
	}
}

func (z ZiplineTemplate) Post(i interface{}, p ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // tempory fix

		handler, err := z.Resolve()
		if err != nil {
			panic(err)
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// write error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (z ZiplineTemplate) Get(i interface{}, p ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // tempory fix

		handler, err := z.Resolve()
		if err != nil {
			panic(err)
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// write error response
			panic(err)
		}
	}
}
