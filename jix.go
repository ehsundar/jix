package jix

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

type Handler[Req, Resp any] func(context.Context, *Req) (*Resp, error)

type Jixer[Req, Resp any] struct {
	handler           Handler[Req, Resp]
	statusMapper      map[error]int
	requestExtractors []RequestExtractor[Req]

	fillRequestHeaders  bool
	fillResponseHeaders bool
	fillQueries         bool
}

func Jixed[Req, Resp any](handler Handler[Req, Resp]) *Jixer[Req, Resp] {
	j := &Jixer[Req, Resp]{
		handler:           handler,
		statusMapper:      make(map[error]int),
		requestExtractors: make([]RequestExtractor[Req], 0),
	}

	j.WithRequestExtractors(BodyExtractor[Req])
	j.WithErrorToStatusMapping(errorToStatusMap)

	return j
}

func (j *Jixer[Req, Resp]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req Req

	ctx = context.WithValue(ctx, "headers", r.Header)
	ctx = context.WithValue(ctx, "queries", r.URL.Query())

	for i, ext := range j.requestExtractors {
		newRequest, err, statusCode := ext(r, req)
		if err != nil {
			logrus.WithError(err).Errorf("extractor failed: %d", i)
			http.Error(w, err.Error(), statusCode)
			return
		}
		req = newRequest
	}

	resp, err := j.handler(ctx, &req)
	if err != nil {
		status := j.getStatusFromError(err)
		http.Error(w, err.Error(), status)
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

func (j *Jixer[Req, Resp]) getStatusFromError(err error) int {
	status, ok := j.statusMapper[err]
	if ok {
		return status
	}

	for e, code := range j.statusMapper {
		if errors.Is(err, e) {
			return code
		}
	}

	logrus.Warnf("no status specified in jixed handler error mapping for: %s", err)

	return http.StatusInternalServerError
}
