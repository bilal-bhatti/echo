package services

import (
	"context"

	"github.com/bilal-bhatti/echo/pkg/connectors"
	"github.com/bilal-bhatti/echo/pkg/models"
)

type ThingsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ThingsService) Create(ctx context.Context, thingRequest *models.ThingRequest) (*models.ThingResponse, error) {
	res := cs.Data.Create(thingRequest.Input)

	return &models.ThingResponse{Output: res}, nil
}

func (cs ThingsService) GetOne(ctx context.Context, id int) (*models.ThingResponse, error) {
	res := cs.Data.GetOne(id)

	return &models.ThingResponse{Output: res}, nil
}
