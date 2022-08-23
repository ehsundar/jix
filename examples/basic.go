package main

import (
	"context"
	"jix"
	"net/http"
)

func main() {
	jixed := jix.Jixed(handler).
		WithFillRequestFromHeader(true).
		WithFillRequestFromQuery(true)
	http.ListenAndServe(":8080", jixed)
}

func handler(ctx context.Context, req *Request) (*Response, error) {
	print(req.AuthToken)
	print(req.SortBy)
	return &Response{Message: "Hello, world!"}, nil
}

type Request struct {
	AuthToken string `json:"auth_token" jix-header:"Authorization"`
	Name      string `json:"name"`
	SortBy    string `json:"sort_by" jix-query:"sort"`
}

type Response struct {
	Message string `json:"message"`
}
