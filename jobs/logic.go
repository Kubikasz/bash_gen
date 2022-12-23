package jobs

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) SortJobs(ctx context.Context, jobs Job) (Job, error) {
	logger := log.With(s.logger, "method", "SortJobs")
	logger.Log("msg", "sorting jobs")
	return jobs, errors.New("not implemented")
}

func (s *service) SortJobsToBash(ctx context.Context, jobs Job) (string, error) {
	logger := log.With(s.logger, "method", "SortJobsToBash")
	logger.Log("msg", "sorting jobs to bash")
	return "", errors.New("not implemented")

}
