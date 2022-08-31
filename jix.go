package jix

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

type Handler[Req, Resp any] func(context.Context, *Req) (*Resp, error)

type Jixer[Req, Resp any] struct {
	handler Handler[Req, Resp]

	fillRequestHeaders  bool
	fillResponseHeaders bool
	fillQueries         bool
}

func Jixed[Req, Resp any](handler Handler[Req, Resp]) *Jixer[Req, Resp] {
	return &Jixer[Req, Resp]{handler: handler}
}

func (j *Jixer[Req, Resp]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req Req

	ctx = context.WithValue(ctx, "headers", r.Header)
	ctx = context.WithValue(ctx, "queries", r.URL.Query())

	j.fillRequestFromHeader(&req, r.Header)
	j.fillRequestFromQueryParams(&req, r.URL.Query())

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := j.handler(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j.fillHeadersFromResponse(resp, w.Header())

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (j *Jixer[Req, Resp]) WithFillRequestFromHeader(fill bool) *Jixer[Req, Resp] {
	j.fillRequestHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillHeadersFromResponse(fill bool) *Jixer[Req, Resp] {
	j.fillResponseHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillRequestFromQuery(fill bool) *Jixer[Req, Resp] {
	j.fillQueries = fill
	return j
}

func (j *Jixer[Req, Resp]) fillRequestFromHeader(r *Req, headers http.Header) {
	if !j.fillRequestHeaders {
		return
	}

	recType := reflect.TypeOf(*r)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-header")
		if tag != "" {
			if val := headers.Get(tag); val != "" {
				reflect.ValueOf(r).Elem().Field(i).SetString(val)
			}
		}
	}
}

func (j *Jixer[Req, Resp]) fillHeadersFromResponse(r *Resp, headers http.Header) {
	if !j.fillResponseHeaders {
		return
	}

	recType := reflect.TypeOf(*r)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-header")
		if tag != "" {
			value := reflect.ValueOf(r).Elem().Field(i).String()
			headers.Set(tag, value)
		}
	}
}

func (j *Jixer[Req, Resp]) fillRequestFromQueryParams(r *Req, queryParams url.Values) {
	if !j.fillQueries {
		return
	}

	recType := reflect.TypeOf(*r)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-query")
		if tag != "" {
			if val := queryParams.Get(tag); val != "" {
				reflect.ValueOf(r).Elem().Field(i).SetString(val)
			}
		}
	}
}
