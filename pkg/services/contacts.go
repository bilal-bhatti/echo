package services

import (
	"context"

	"github.com/bilal-bhatti/echo/pkg/connectors"
)

type ContactRequest struct {
	Input string `json:"input"`
}

type ContactResponse struct {
	Output string         `json:"output"`
	Input  ContactRequest `json:"input"`
}

type ContactsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ContactsService) Create(ctx context.Context, contactRequest *ContactRequest) (*ContactResponse, error) {
	res := cs.Data.Create(contactRequest.Input)

	return &ContactResponse{Output: res, Input: *contactRequest}, nil
}

func (cs ContactsService) GetOne(id int) (*ContactResponse, error) {
	res := cs.Data.GetOne(id)

	return &ContactResponse{Output: res}, nil
}
