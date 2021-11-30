//go:build wireinject
// +build wireinject

package services

import (
	"context"
	"net/http"

	"github.com/bilal-bhatti/echo/pkg/connectors"

	"github.com/google/wire"
)

func ProvideContext(r *http.Request) context.Context {
	return r.Context()
}

func InitContactsService() (*ContactsService, error) {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, wire.Struct(new(ContactsService), "*")))
}

func InitThingsService() (*ThingsService, error) {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, wire.Struct(new(ThingsService), "*")))
}
