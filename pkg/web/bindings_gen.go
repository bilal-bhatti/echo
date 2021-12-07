// Code generated by Zipline. DO NOT EDIT.

//go:build !ziplinegen
// +build !ziplinegen

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

	mux.Post("/contacts", ContactsServiceCreateHandlerFunc())
	mux.Get("/contacts/{id}", ContactsServiceGetOneHandlerFunc())

	mux.Post("/things", ThingsServiceCreateHandlerFunc())
	mux.Get("/things/{id}", ThingsServiceGetOneHandlerFunc())

	return mux
}

// ContactsServiceCreateHandlerFunc handles requests to:
// path  : /contacts
// method: post
func ContactsServiceCreateHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		// initialize application handler
		handler, err := services.InitContactsService()
		if err != nil {

			skit.Failure(w, err)
			return
		}

		// resolve parameter [ctx] through a provider
		ctx := services.ProvideContext(r)

		// resolve parameter [contactRequest] with [Body] template
		defer io.Copy(ioutil.Discard, r.Body)
		contactRequest := &services.ContactRequest{}
		err = json.NewDecoder(r.Body).Decode(contactRequest)
		if err != nil {
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to decode request body"))
			return
		}

		// execute application handler
		response, err := handler.Create(ctx, contactRequest)
		if err != nil {
			skit.Failure(w, err)
			return
		}
		skit.Success(w, response)
	}
}

// ContactsServiceGetOneHandlerFunc handles requests to:
// path  : /contacts/{id}
// method: get
func ContactsServiceGetOneHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		// initialize application handler
		handler, err := services.InitContactsService()
		if err != nil {

			skit.Failure(w, err)
			return
		}

		// resolve parameter [ctx] through a provider
		ctx := services.ProvideContext(r)

		// resolve parameter [id] with [Path] template
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to parse path parameter as int"))
			return
		}

		// execute application handler
		response, err := handler.GetOne(ctx, id)
		if err != nil {

			skit.Failure(w, err)
			return
		}
		skit.Success(w, response)
	}
}

// ThingsServiceCreateHandlerFunc handles requests to:
// path  : /things
// method: post
func ThingsServiceCreateHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		// initialize application handler
		handler, err := services.InitThingsService()
		if err != nil {

			skit.Failure(w, err)
			return
		}

		// resolve parameter [ctx] through a provider
		ctx := services.ProvideContext(r)

		// resolve parameter [thingRequest] with [Body] template
		defer io.Copy(ioutil.Discard, r.Body)
		thingRequest := &services.ThingRequest{}
		err = json.NewDecoder(r.Body).Decode(thingRequest)
		if err != nil {
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to decode request body"))
			return
		}

		// execute application handler
		response, err := handler.Create(ctx, thingRequest)
		if err != nil {
			skit.Failure(w, err)
			return
		}
		skit.Success(w, response)
	}
}

// ThingsServiceGetOneHandlerFunc handles requests to:
// path  : /things/{id}
// method: get
func ThingsServiceGetOneHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error // why not

		startTime := time.Now()
		defer func() {
			duration := time.Now().Sub(startTime)
			log.Printf("It took %s to process request\n", duration.String())
		}()

		// initialize application handler
		handler, err := services.InitThingsService()
		if err != nil {

			skit.Failure(w, err)
			return
		}

		// resolve parameter [ctx] through a provider
		ctx := services.ProvideContext(r)

		// resolve parameter [id] with [Path] template
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			skit.Failure(w, skit.WithStatus(err, http.StatusBadRequest, "failed to parse path parameter as int"))
			return
		}

		// execute application handler
		response, err := handler.GetOne(ctx, id)
		if err != nil {

			skit.Failure(w, err)
			return
		}
		skit.Success(w, response)
	}
}
