// Code generated by Zipline. DO NOT EDIT.

//go:generate zipline
// +build !ziplinegen

package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bilal-bhatti/echo/pkg/models"
	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Post("/contacts", ContactsServiceCreateHandlerFunc())
	mux.Get("/contacts/{id}", ContactsServiceGetOneHandlerFunc())
	mux.Post("/things", ThingsServiceCreateHandlerFunc())
	mux.Get("/things/{id}", ThingsServiceGetOneHandlerFunc())
	mux.Post("/echo", EchoHandlerFunc())
	return mux
}

// ContactsServiceCreateHandlerFunc handles requests to:
// path  : /contacts
// method: post
func ContactsServiceCreateHandlerFunc() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		// resolve ctx dependency through a provider function
		ctx := ProvideContext(request)

		// extract json body and marshall contactRequest
		contactRequest := &models.ContactRequest{}
		err = json.NewDecoder(request.Body).Decode(contactRequest)
		if err != nil {
			// invalid request error
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// initialize application handler
		contactsService := InitContactsService()
		// execute application handler
		response, err := contactsService.Create(ctx, contactRequest)
		if err != nil {
			panic(err)
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

// ContactsServiceGetOneHandlerFunc handles requests to:
// path  : /contacts/{id}
// method: get
func ContactsServiceGetOneHandlerFunc() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		// parse path parameter id
		id, err := strconv.Atoi(chi.URLParam(request, "id"))
		if err != nil {
			// invalid request error
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// initialize application handler
		contactsService := InitContactsService()
		// execute application handler
		response, err := contactsService.GetOne(id)
		if err != nil {
			panic(err)
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

// ThingsServiceCreateHandlerFunc handles requests to:
// path  : /things
// method: post
func ThingsServiceCreateHandlerFunc() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		// extract json body and marshall thingRequest
		thingRequest := &models.ThingRequest{}
		err = json.NewDecoder(request.Body).Decode(thingRequest)
		if err != nil {
			// invalid request error
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// initialize application handler
		thingsService := InitThingsService()
		// execute application handler
		response, err := thingsService.Create(thingRequest)
		if err != nil {
			panic(err)
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

// ThingsServiceGetOneHandlerFunc handles requests to:
// path  : /things/{id}
// method: get
func ThingsServiceGetOneHandlerFunc() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		// parse path parameter id
		id, err := strconv.Atoi(chi.URLParam(request, "id"))
		if err != nil {
			// invalid request error
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// initialize application handler
		thingsService := InitThingsService()
		// execute application handler
		response, err := thingsService.GetOne(id)
		if err != nil {
			panic(err)
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}

// EchoHandlerFunc handles requests to:
// path  : /echo
// method: post
func EchoHandlerFunc() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var err error // tempory fix

		// extract json body and marshall echoRequest
		echoRequest := EchoRequest{}
		err = json.NewDecoder(request.Body).Decode(&echoRequest)
		if err != nil {
			// invalid request error
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// execute application handler
		response, err := Echo(echoRequest)
		if err != nil {
			panic(err)
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(responseWriter).Encode(response)
		if err != nil {
			panic(err)
		}
	}
}
