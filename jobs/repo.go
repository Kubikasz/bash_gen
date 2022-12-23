package jobs

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to handle Repo request")

// used to have a starting point for db implementation
// add db entity here
type repo struct {
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		logger: logger,
	}
}

func (r *repo) CreateTask(ctx context.Context, task Task) error {
	logger := log.With(r.logger, "method", "CreateTask")
	logger.Log("msg", "creating task")
	return nil
}

func (r *repo) GetTask(ctx context.Context, name string) (Task, error) {
	logger := log.With(r.logger, "method", "GetTask")
	logger.Log("msg", "getting task")
	return Task{}, nil
}
