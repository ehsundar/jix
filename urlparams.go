package jix

import (
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
