package main

import (
	"context"
	"errors"
	"fmt"
	"jix"
	"net/http"
)

var (
	ErrNoNameProvided = errors.New("no name provided")
)

func main() {
	jixed := jix.Jixed(handler).
		WithFillRequestFromHeader(true).
		WithFillRequestFromQuery(true).
		WithFillHeadersFromResponse(true).
		WithErrorToStatusMapping(map[error]int{
			ErrNoNameProvided: 400,
		})
	http.ListenAndServe(":8080", jixed)
}

func handler(ctx context.Context, req *Request) (*Response, error) {
	print(req.AuthToken)
	print(req.SortBy)
	if len(req.Name) == 0 {
		return nil, fmt.Errorf("%s: %w", ErrNoNameProvided, jix.ErrAborted)
	}
	return &Response{Message: "Hello, world!", ValidUntil: "2023"}, nil
}

type Request struct {
	AuthToken string `json:"auth_token" jix-header:"Authorization"`
	Name      string `json:"name"`
	SortBy    string `json:"sort_by" jix-query:"sort"`
}

type Response struct {
	Message    string `json:"message"`
	ValidUntil string `json:"-" jix-header:"X-Valid-Until"`
}
