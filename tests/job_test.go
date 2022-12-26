package tests

// testing the logic only of the serive
// read json data from file and parse it
import (
	"bash_gen/jobs"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestSortJobs(t *testing.T) {
	// create a job service

	logger := log.NewLogfmtLogger(os.Stderr)
	repo := jobs.NewRepo(logger)
	srv := jobs.NewService(repo, logger)
	ctx := context.Background()

	// setup table test cases good and invalid

	job_ok_req, err := loadJobFromFile("testdata/job_good_req.json", t)
	if err != nil {
		t.Fatal(err)
	}

	job_ok_resp := jobs.Job{}
	file, _ := os.Open("testdata/job_goo_resp.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&job_ok_resp)
	if err != nil {
		t.Errorf("error decoding json: %v", err)
	}

	job_circ_req, err := loadJobFromFile("testdata/job_circular_req.json", t)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(job_circ_req)

	job_invalid, err := loadJobFromFile("testdata/job_invalid_req.json", t)
	if err != nil {
		t.Fatal(err)
	}

	type test struct {
		name     string
		job      jobs.Job
		expected jobs.Job
		err      error
	}
	tests := []test{
		{name: "ok",
			job:      job_ok_req,
			expected: job_ok_resp,
			err:      nil,
		},
		{name: "circular",
			job:      job_circ_req,
			expected: nil,
			err:      jobs.ErrCantSolveGraph,
		},
		{
			name:     "invalid",
			job:      job_invalid,
			expected: nil,
			err:      jobs.ErrInvalidJob,
		},
	}

	// run the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			resp, err := srv.SortJobs(ctx, tc.job)
			if err != nil {
				if tc.err != nil && err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
				return

			}
			if len(resp) != len(tc.expected) {
				t.Errorf("expected response with len:  %v, got %v", len(tc.expected), len(resp))
				return
			}
			for i := range tc.expected {
				if resp[i].Name != tc.expected[i].Name {
					t.Errorf("expected response %v, got %v", tc.expected, resp)
				}
			}
		})
	}

}

func loadJobFromFile(filename string, t *testing.T) (jobs.Job, error) {
	// read json data from testdata/job_good_req.json
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	return jsonToJob(file)
}

func jsonToJob(j io.Reader) (jobs.Job, error) {
	type Tasks struct {
		Tasks jobs.Job `json:"tasks"`
	}
	var tasks Tasks
	decoder := json.NewDecoder(j)
	err := decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks.Tasks, nil
}
