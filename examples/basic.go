package main

import (
	"context"
	"jix"
	"net/http"
)

func main() {
	jixed := jix.Jixed(handler).WithFillRequestFromHeader(true)
	http.ListenAndServe(":8080", jixed)
}

func handler(ctx context.Context, req *Request) (*Response, error) {
	print(req.AuthToken)
	return &Response{Message: "Hello, world!"}, nil
}

type Request struct {
	AuthToken string `json:"auth_token" jix-header:"Authorization"`
	Name      string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}
