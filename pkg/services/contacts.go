package services

import (
	"context"
	"github.com/bilal-bhatti/echo/pkg/connectors"
	"github.com/bilal-bhatti/echo/pkg/models"
	"log"
)

type ContactsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ContactsService) Create(ctx context.Context, contactRequest *models.ContactRequest) (*models.ContactResponse, error) {
	res := cs.Data.Create(contactRequest.Input)

	return &models.ContactResponse{Output: res}, nil
}

func (cs ContactsService) GetOne(id int) (*models.ContactResponse, error) {
	res := cs.Data.GetOne(id)

	x := "x"
	log.Println("x", x)

	return &models.ContactResponse{Output: res}, nil
}
