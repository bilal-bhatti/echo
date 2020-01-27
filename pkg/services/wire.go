// +build wireinject

package services

import (
	"context"
	"echo/pkg/connectors"
	"net/http"

	"github.com/google/wire"
)

func ProvideContext(r *http.Request) context.Context {
	return r.Context()
}

func InitContactsService() *ContactsService {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, ContactsService{}))
}

func InitThingsService() *ThingsService {
	panic(wire.Build(connectors.ProvideDataConnector, connectors.ProvideIdentityConnector, connectors.ProvideElasticSearchConnector, ThingsService{}))
}
