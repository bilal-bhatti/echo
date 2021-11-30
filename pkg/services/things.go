package services

import (
	"context"

	"github.com/bilal-bhatti/echo/pkg/connectors"
)

type ThingRequest struct {
	Input string `json:"input"`
}

type ThingResponse struct {
	Output string       `json:"output"`
	Input  ThingRequest `json:"input"`
}

type ThingsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ThingsService) Create(thingRequest *ThingRequest) (*ThingResponse, error) {
	res := cs.Data.Create(thingRequest.Input)

	return &ThingResponse{Output: res, Input: *thingRequest}, nil
}

func (cs ThingsService) GetOne(id int) (*ThingResponse, error) {
	res := cs.Data.GetOne(id)

	return &ThingResponse{Output: res}, nil
}
