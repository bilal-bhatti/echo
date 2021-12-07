//go:build ziplinegen
// +build ziplinegen

package web

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bilal-bhatti/echo/pkg/services"
	"github.com/bilal-bhatti/skit"
	"github.com/go-chi/chi/v5"
)

func NewRouter(mux *chi.Mux) *chi.Mux {
	// bind contacts
	mux.Post("/contacts", z.Post(new(services.ContactsService).Create, z.Resolve, z.Body))
	mux.Get("/contacts/{id}", z.Get(new(services.ContactsService).GetOne, z.Resolve, z.Path))

	// bind things
	mux.Post("/things", z.Post(services.ThingsService.Create, z.Resolve, z.Body))
	mux.Get("/things/{id}", z.Get(services.ThingsService.GetOne, z.Resolve, z.Path))

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
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to parse path parameter as int"))
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
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to parse query paremeter as int"))
			return
		}
		z.DevNull(name)
	}
}

func (z ZiplineTemplate) Body(w http.ResponseWriter, r *http.Request) {
	var err error
	defer io.Copy(ioutil.Discard, r.Body)
	name := &ZiplineTemplate{}
	err = json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to decode request body"))
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
			skit.Failure(w, err)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			skit.Failure(w, err)
			return
		}

		skit.Success(w, response)
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
			skit.Failure(w, err)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			// write error response
			skit.Failure(w, err)
			return
		}

		skit.Success(w, response)
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
			skit.Failure(w, err)
			return
		}

		err = handler.ReturnError()

		if err != nil {
			// write error response
			skit.Failure(w, err)
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
			skit.Failure(w, err)
			return
		}

		response, err := handler.ReturnResponseAndError()
		if err != nil {
			// write error response
			skit.Failure(w, err)
			return
		}

		skit.Success(w, response)
	}
}
