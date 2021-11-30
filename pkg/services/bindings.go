//go:build ziplinegen
// +build ziplinegen

package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func NewRouter(mux *chi.Mux) *chi.Mux {
	// bind contacts
	mux.Post("/contacts", z.Post(new(ContactsService).Create, z.Resolve, z.Body))
	mux.Get("/contacts/{id}", z.Get(new(ContactsService).GetOne, z.Resolve, z.Path))

	// bind things
	mux.Post("/things", z.Post(ThingsService.Create, z.Resolve, z.Body))
	mux.Get("/things/{id}", z.Get(ThingsService.GetOne, z.Resolve, z.Path))

	//mux.Get("/echo/{str}", z.Get(services.Echo, z.Resolve, z.Path))

	return mux
}

var z ZiplineTemplate

type ZiplineTemplate struct {
	ReturnResponseAndError func() (interface{}, error)
	ReturnError            func() error
	Resolve                func() (ZiplineTemplate, error)
	DevNull                func(i ...interface{})
}

func (z ZiplineTemplate) Path(kind string, w http.ResponseWriter, r *http.Request) {
	switch kind {
	case "string":
		name := chi.URLParam(r, "name")
		z.DevNull(name)
	case "int":
		name, err := strconv.Atoi(chi.URLParam(r, "name"))
		if err != nil {
			// invalid request error
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		z.DevNull(name)
	}
}

func (z ZiplineTemplate) Query(kind string, w http.ResponseWriter, r *http.Request) {
	switch kind {
	case "string":
		name := r.URL.Query().Get("name")
		z.DevNull(name)
	case "int":
		name, err := strconv.Atoi(r.URL.Query().Get("name"))
		if err != nil {
			// invalid request error
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		z.DevNull(name)
	}
}

func (z ZiplineTemplate) Body(w http.ResponseWriter, r *http.Request) {
	var err error
	name := &ZiplineTemplate{}
	err = json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		// invalid request error
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func (z ZiplineTemplate) Post(i interface{}, p ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		handler, err := z.Resolve()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (z ZiplineTemplate) Get(i interface{}, p ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		handler, err := z.Resolve()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

func (z ZiplineTemplate) Delete(i interface{}, params ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		handler, err := z.Resolve()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = handler.ReturnError()

		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (z ZiplineTemplate) Put(i interface{}, p ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		handler, err := z.Resolve()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			// write error response
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
