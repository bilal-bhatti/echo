package services

import (
	"echo/pkg/connectors"
	"echo/pkg/models"
)

type ThingsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ThingsService) Create(req *models.ThingRequest) (*models.ThingResponse, error) {
	res := cs.Data.Create(req.Input)

	return &models.ThingResponse{Output: res}, nil
}

func (cs ThingsService) GetOne(id int) (*models.ThingResponse, error) {
	res := cs.Data.GetOne(id)

	return &models.ThingResponse{Output: res}, nil
}
