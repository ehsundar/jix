package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"jix"
	"net/http"
)

var (
	ErrNoNameProvided = errors.New("no name provided")
)

type Request struct {
	AuthToken string `json:"auth_token" jix-header:"Authorization"`
	Name      string `json:"name"`
	SortBy    string `json:"sort_by" jix-query:"sort"`
	Category  string `json:"-" jix-param:"category"`
}

type Response struct {
	Message    string `json:"message"`
	Target     string `json:"target"`
	ValidUntil string `json:"-" jix-header:"X-Valid-Until"`
}

func handler(ctx context.Context, req *Request) (*Response, error) {
	print(req.AuthToken)
	print(req.SortBy)
	if len(req.Name) == 0 {
		return nil, fmt.Errorf("%s: %w", ErrNoNameProvided, jix.ErrAborted)
	}
	return &Response{
		Message:    "Hello, world!",
		ValidUntil: "2023",
		Target:     req.Category,
	}, nil
}

func main() {
	jixed := jix.Jixed(handler).
		WithFillRequestFromHeader(true).
		WithFillRequestFromQuery(true).
		WithFillHeadersFromResponse(true).
		WithErrorToStatusMapping(map[error]int{
			ErrNoNameProvided: 400,
		}).
		WithRequestExtractors(jix.GorillaMuxURLParamsExtractor[Request])

	router := mux.NewRouter()
	router.Handle("/example/{category}/basix/", jixed)

	http.ListenAndServe(":8080", router)
}
