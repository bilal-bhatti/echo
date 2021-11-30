package services

import (
	"context"
	"strings"
)

type EchoResponse struct {
	Output string `json:"output"`
}

func Echo(ctx context.Context, str string) (EchoResponse, error) {
	return EchoResponse{
		Output: strings.Replace(str, "i", "o", -1),
	}, nil
}
