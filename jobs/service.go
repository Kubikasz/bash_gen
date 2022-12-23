package jobs

import "context"

type Service interface {
	// sort jobs context and jobs returns a sorted list of jobs and error
	SortJobs(ctx context.Context, jobs Job) (Job, error)
	SortJobsToBash(ctx context.Context, jobs Job) (string, error)
}
