package jobs

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SortJobsEndpoint       endpoint.Endpoint
	SortJobsToBashEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		SortJobsEndpoint:       makeSortJobsEndpoint(s),
		SortJobsToBashEndpoint: makeSortJobsToBashEndpoint(s),
	}
}

func makeSortJobsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Job)
		jobs, err := s.SortJobs(ctx, req)
		if err != nil {
			return nil, err
		}
		return jobs, nil
	}
}

func makeSortJobsToBashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Job)
		bash, err := s.SortJobsToBash(ctx, req)
		if err != nil {
			return nil, err
		}
		return bash, nil
	}
}
