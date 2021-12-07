package services

import (
	"context"
	"net/http"

	"github.com/bilal-bhatti/echo/pkg/connectors"
	"github.com/bilal-bhatti/skit"
	"github.com/pkg/errors"
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
	//	res := cs.Data.Create(contactRequest.Input)
	// 400 Bad Request – client sent an invalid request, such as lacking required request body or parameter
	// - request parsing time
	// - business rules validation
	// 401 Unauthorized – client failed to authenticate with the server
	// 403 Forbidden – client authenticated but does not have permission to access the requested resource
	// 404 Not Found – the requested resource does not exist
	// 412 Precondition Failed – one or more conditions in the request header fields evaluated to false
	// 500 Internal Server Error – a generic error occurred on the server
	// 503 Service Unavailable – the requested service is not available

	//return &ContactResponse{Output: res, Input: *contactRequest}, nil
	return nil, skit.WithStatus(errors.New("something went wrong"), http.StatusBadRequest, []string{"one", "two"})
}

func (cs ContactsService) GetOne(ctx context.Context, id int) (*ContactResponse, error) {
	res := cs.Data.GetOne(id)

	return &ContactResponse{Output: res}, nil
}
