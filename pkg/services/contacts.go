package services

import (
	"context"
	"echo/pkg/connectors"
	"echo/pkg/models"
	"log"
)

type ContactsService struct {
	Data          *connectors.DataConnector
	Identity      *connectors.IdentityConnector
	ElasticSearch *connectors.ElasticSearchConnector
}

func (cs ContactsService) Create(ctx context.Context, req *models.ContactRequest) (*models.ContactResponse, error) {
	res := cs.Data.Create(req.Input)

	return &models.ContactResponse{Output: res}, nil
}

func (cs ContactsService) GetOne(id int) (*models.ContactResponse, error) {
	res := cs.Data.GetOne(id)

	x := "x"
	log.Println("x", x)

	return &models.ContactResponse{Output: res}, nil
}

func (cs ContactsService) Test() {}

func Test() {
	spec := ContactsService{}
	spec.Test()
	log.Println("spec", spec)
}
