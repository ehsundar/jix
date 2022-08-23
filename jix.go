package jix

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

type Jixer[Req, Resp any] struct {
	handler     func(context.Context, *Req) (*Resp, error)
	fillHeaders bool
	fillQueries bool
}

func Jixed[Req, Resp any](handler func(context.Context, *Req) (*Resp, error)) *Jixer[Req, Resp] {
	return &Jixer[Req, Resp]{handler: handler}
}

func (j *Jixer[Req, Resp]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req Req

	if j.fillHeaders {
		ctx = context.WithValue(ctx, "headers", r.Header)
		j.fillRequestFromHeader(&req, r.Header)
	}
	if j.fillQueries {
		ctx = context.WithValue(ctx, "queries", r.URL.Query())
		j.fillRequestFromQueryParams(&req, r.URL.Query())
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := j.handler(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (j *Jixer[Req, Resp]) WithFillRequestFromHeader(fill bool) *Jixer[Req, Resp] {
	j.fillHeaders = fill
	return j
}

func (j *Jixer[Req, Resp]) WithFillRequestFromQuery(fill bool) *Jixer[Req, Resp] {
	j.fillQueries = fill
	return j
}

func (j *Jixer[Req, Resp]) fillRequestFromHeader(r *Req, headers http.Header) {
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

func (j *Jixer[Req, Resp]) fillRequestFromQueryParams(r *Req, queryParams url.Values) {
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
