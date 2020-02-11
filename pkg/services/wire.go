// +build wireinject

package services

import (
	"context"
	"github.com/bilal-bhatti/echo/pkg/connectors"
	"net/http"

	"github.com/google/wire"
)

func ProvideContext(r *http.Request) context.Context {
	return r.Context()
}

func InitContactsService() *ContactsService {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, wire.Struct(new(ContactsService), "*")))
}

func InitThingsService() *ThingsService {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, wire.Struct(new(ThingsService), "*")))
}
