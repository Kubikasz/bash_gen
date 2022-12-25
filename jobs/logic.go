package jobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/stevenle/topsort"
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
	// log all the sorted jobs
	for _, task := range sortedJobs {
		logger.Log("task", task.Name)
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

// jobToGraphto convert job to list of strings as edges
func jobToGraphto(job Job) ([]string, error) {
	// for each task in the job add the task name as node and the dependencies as edges
	graph := topsort.NewGraph()
	noDepNodes := []string{}
	for _, task := range job {
		graph.AddNode(task.Name)
		if task.Requires == nil {
			// no dependencies
			noDepNodes = append(noDepNodes, task.Name)
			continue
		}
		for _, dep := range *task.Requires {
			graph.AddEdge(task.Name, dep)
		}
	}
	if len(noDepNodes) == 0 {
		return nil, ErrCantSolveGraph
	}
	// sort the graph
	sorted, err := graph.TopSort(noDepNodes[0])
	if err != nil {
		return nil, err
	}
	return sorted, nil

}

// treeToJov gets list of names as order and orders the job by these names
func treeToJob(sorted []string, job Job) (Job, error) {
	// for each sorted name find the task and add it to the new job
	if len(sorted) == 0 {
		return nil, ErrCantSolveGraph
	}
	sortedJob := Job{}
	for _, name := range sorted {
		for _, task := range job {
			if task.Name == name {
				// set requires to nil
				fmt.Println(task.Name)
				task.Requires = nil
				sortedJob = append(sortedJob, task)
			}
		}
	}
	return sortedJob, nil
}

// sortJobs to sort the job and return sorted job
func sortJobs(job Job) (Job, error) {
	sortedTasks, err := jobToGraphto(job)
	if err != nil {
		return nil, err
	}
	sortedJob, err := treeToJob(sortedTasks, job)
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
