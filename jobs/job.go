package jobs

import (
	"context"
	"fmt"
)

// Task represents the task to be excuted and its dependencies
// using pointer to string to be able to avoid when responding
type Task struct {
	Name     string    `json:"name"`
	Command  string    `json:"command"`
	Requires *[]string `json:"requires,omitempty"`
	ID       *int      `json:"id,omitempty"`
}

type Job []Task

// used to have a starting point for db implementation
type Repository interface {
	CreateTask(ctx context.Context, task Task) error
	GetTask(ctx context.Context, name string) (Task, error)
}

// implemnt String method for Task
func (t Task) String() string {
	id := ""
	if t.ID != nil {
		id = fmt.Sprintf("%d, ", *t.ID)
	}
	return fmt.Sprint("Task: ", t.Name, " Command: ", t.Command, " Requires: ", t.Requires, " ID: ", id)
}

// add validation to Task
func (t Task) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("name is required")
	}
	if t.Command == "" {
		return fmt.Errorf("command is required")
	}
	return nil
}

// add validation to Job
func (j Job) Validate() error {
	for _, task := range j {
		if err := task.Validate(); err != nil {
			return err
		}
	}
	return nil
}
