package jobs

import (
	"context"
	"errors"
	"net/http"

	"encoding/json"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/sort").Handler(httptransport.NewServer(
		endpoints.SortJobsEndpoint,
		decodeSortJobsRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/sort/bash").Handler(httptransport.NewServer(
		endpoints.SortJobsToBashEndpoint,
		decodeSortJobsToBashRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func decodeSortJobsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Job

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("error decoding request")
	}
	return req, nil
}

func decodeSortJobsToBashRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req Job
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
