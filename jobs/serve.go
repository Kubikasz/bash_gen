package jobs

import (
	"context"
	"net/http"
  "fmt"

	"encoding/json"
	"io"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/sort/bash").Handler(httptransport.NewServer(
		endpoints.SortJobsToBashEndpoint,
		decodeSortJobsToBashRequest,
		encodeBashResponse,
	))
	r.Use(commonMiddleware)
	r.Methods("POST").Path("/sort").Handler(httptransport.NewServer(
		endpoints.SortJobsEndpoint,
		decodeSortJobsRequest,
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
	return JsonToJob(r.Body)
}

func decodeSortJobsToBashRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return JsonToJob(r.Body)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeBashResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
  resp, ok := response.(string)
  if !ok {
    return fmt.Errorf("response is not a string")
  }
  w.Write([]byte(resp))
  return nil
}

func JsonToJob(j io.Reader) (Job, error) {
	type Tasks struct {
		Tasks Job `json:"tasks"`
	}
	var tasks Tasks
	decoder := json.NewDecoder(j)
	err := decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks.Tasks, nil
}
