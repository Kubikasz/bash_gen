package tests

// testing the logic only of the serive
// read json data from file and parse it
import (
	"bash_gen/jobs"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"os"
	"testing"
)

func TestSortJobs(t *testing.T) {
	// create a job service

	logger := log.NewLogfmtLogger(os.Stderr)
	repo := jobs.NewRepo(logger)
	srv := jobs.NewService(repo, logger)
	ctx := context.Background()

	// setup table test cases good and invalid

	type Tasks struct {
		Tasks jobs.Job `json:"tasks"`
	}
	var tasks Tasks
	// read json data from testdata/job_good_req.json
	file, _ := os.Open("testdata/job_good_req.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&tasks)
	if err != nil {
		t.Errorf("error decoding json: %v", err)
	}
	job_ok_req := tasks.Tasks

	// job_invalid := jobs.Job{}

	type test struct {
		name     string
		job      jobs.Job
		expected jobs.Job
		err      error
	}
	tests := []test{
		{name: "ok",
			job:      job_ok_req,
			expected: job_ok_req,
			err:      nil,
		},
		// {
		// 	name:     "invalid",
		// 	job:      &job_invalid,
		// 	expected: nil,
		// 	err:      errors.New("invalid job"),
		// },
	}

	// run the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := srv.SortJobs(ctx, tc.job)
			if err != nil {
				if tc.err != nil && err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
			}
			if resp == nil {
				t.Errorf("expected response %v, got %v", tc.expected, resp)
			}

			for i := range resp {
				if resp[i].Name != tc.expected[i].Name {
					t.Errorf("expected response %v, got %v", tc.expected, resp)
				}
			}
		})
	}

}
