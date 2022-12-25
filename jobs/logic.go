package jobs

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/yourbasic/graph"
)

// create error message from job
var ErrInvalidJob = errors.New("invalid job")
var ErrCycle = errors.New("cycle detected")
var ErrCantSolveGraph = errors.New("can't solve graph")

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
	sortedJobs, err := sortJobs(jobs)
	if err != nil {
		logger.Log("err", err)
		return nil, err
	}

	return sortedJobs, nil
}

func (s *service) SortJobsToBash(ctx context.Context, jobs Job) (string, error) {
	logger := log.With(s.logger, "method", "SortJobsToBash")
	logger.Log("msg", "sorting jobs to bash")
	bashString, err := sortJobsToBash(jobs)
	if err != nil {
		logger.Log("err", err)
		return "", err
	}
	return bashString, nil

}

// jobToGraph adds id and returns a graph of the jobs
func jobToGraph(job Job) (j Job, gr *graph.Mutable, e error) {
	// create a graph
	g := graph.New(len(job))
	// add id to each task
	for i, _ := range job {
		var id int = i
		job[i].ID = &id
	}
	// add each task to the graph
	for _, task := range job {
		// add task to the graph
		if task.Requires == nil {
			continue
		}
		// add each dependency to the graph
		for _, require := range *task.Requires {
			for _, r := range job {
				// find the task id of the dependency
				if r.Name == require {
					g.Add(*r.ID, *task.ID)
				}
			}
		}
	}

	// add edges to the graph
	return job, g, nil
}

// treeToJov gets list of names as order and orders the job by these names
func sortedGraphToJob(sorted []int, job Job) (Job, error) {
	sortedJob := Job{}
	for _, id := range sorted {
		for _, task := range job {
			if *task.ID == id {
				sortedJob = append(sortedJob, task)
			}
		}
	}
	return sortedJob, nil
}

// sortJobs to sort the job and return sorted job
func sortJobs(job Job) (Job, error) {
	jobFilled, tasksGraph, err := jobToGraph(job)
	if err != nil {
		return nil, err
	}
	sortedTasks, ok := graph.TopSort(tasksGraph)
	if ok != true {
		return nil, ErrCantSolveGraph
	}
	sortedJob, err := sortedGraphToJob(sortedTasks, jobFilled)
	if err != nil {
		return nil, err
	}

	return sortedJob, nil
}

// sortJobsToBash to sort the job and return bash script as string
func sortJobsToBash(job Job) (string, error) {
	sortedJobs, err := sortJobs(job)
	if err != nil {
		return "", err
	}
	bashString := ""
	for _, task := range sortedJobs {
		// add each Command to the bash with a new line
		bashString += task.Command + "\n"

	}
	return bashString, nil
}
