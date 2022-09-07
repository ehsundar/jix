package jix

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

type RequestExtractor[Req any] func(r *http.Request, req Req) (enrichedReq Req, err error, statusCode int)

func GorillaMuxURLParamsExtractor[Req any](r *http.Request, req Req) (enrichedReq Req, err error, statusCode int) {
	enrichedReq = req

	vars := mux.Vars(r)
	if vars == nil {
		return
	}

	recType := reflect.TypeOf(enrichedReq)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-param")
		if tag != "" {
			if val, ok := vars[tag]; ok {
				reflect.ValueOf(&enrichedReq).Elem().Field(i).SetString(val)
			}
		}
	}

	return
}

func HeaderExtractor[Req any](r *http.Request, req Req) (enrichedReq Req, err error, statusCode int) {
	enrichedReq = req
	headers := r.Header

	recType := reflect.TypeOf(enrichedReq)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-header")
		if tag != "" {
			if val := headers.Get(tag); val != "" {
				reflect.ValueOf(&enrichedReq).Elem().Field(i).SetString(val)
			}
		}
	}

	return
}

func QueryExtractor[Req any](r *http.Request, req Req) (enrichedReq Req, err error, statusCode int) {
	enrichedReq = req
	queryParams := r.URL.Query()

	recType := reflect.TypeOf(enrichedReq)
	for i := 0; i < recType.NumField(); i++ {
		field := recType.Field(i)
		tag := field.Tag.Get("jix-query")
		if tag != "" {
			if val := queryParams.Get(tag); val != "" {
				reflect.ValueOf(&enrichedReq).Elem().Field(i).SetString(val)
			}
		}
	}

	return
}

func BodyExtractor[Req any](r *http.Request, req Req) (enrichedReq Req, err error, statusCode int) {
	enrichedReq = req
	if err = json.NewDecoder(r.Body).Decode(&enrichedReq); err != nil {
		statusCode = http.StatusBadRequest
		return
	}
	return
}
