package services

import (
	"net/http"
	"strings"
)

type EchoRequest struct {
	Input string `json:"input"`
}

type EchoResponse struct {
	Output string `json:"output"`
}

func Echo(req EchoRequest) (EchoResponse, error) {
	return EchoResponse{
		Output: strings.Replace(req.Input, "i", "o", -1),
	}, nil
}

func Zipline(funk func(req EchoRequest) (EchoResponse, error)) http.HandlerFunc {
	panic("hi")
}
