package executor

import (
	"fmt"
	"testing"

	"github.com/Clever/workflow-manager/resources"
	"github.com/Clever/workflow-manager/store"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

type mockBatchClient struct {
	tasks map[string]resources.Task
}

func (be *mockBatchClient) SubmitJob(name string, definition string, dependencies, input []string) (string, error) {
	for _, d := range dependencies {
		if _, ok := be.tasks[d]; !ok {
			return "", fmt.Errorf("Dependency %s not found", d)
		}
	}
	taskID := uuid.NewV4().String()
	be.tasks[taskID] = resources.Task{
		ID:    taskID,
		Name:  name,
		Input: input,
	}

	return taskID, nil
}

func (be *mockBatchClient) Status(tasks []*resources.Task) []error {
	errs := []error{}
	for _, t := range tasks {
		if _, ok := be.tasks[t.ID]; !ok {
			errs = append(errs, fmt.Errorf("%s", t.ID))
		}
	}
	return errs
}

func (be *mockBatchClient) Cancel(tasks []*resources.Task, reason string) []error {
	// mark first task as Cancelled
	if len(tasks) > 0 {
		tasks[0].Status = resources.TaskStatusUserAborted
	}

	return nil
}

func TestUpdateJobStatus(t *testing.T) {
	jm := BatchJobManager{
		&mockBatchClient{
			map[string]resources.Task{},
		},
		store.NewMemoryStore(),
	}
	wf := resources.KitchenSinkWorkflow(t)
	input := []string{"test-start-input"}

	job, err := jm.CreateJob(wf, input)
	assert.Nil(t, err)

	t.Log("Job is QUEUED till a task starts RUNNING")
	assert.Equal(t, job.Status, resources.Queued)

	// mark one task as running
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusRunning)
		break
	}
	err = jm.UpdateJobStatus(job)
	t.Log("Job is RUNNING when a task starts RUNNING")
	assert.Nil(t, err)
	assert.Equal(t, job.Status, resources.Running)

	// mark one task as failed
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusFailed)
		break
	}
	err = jm.UpdateJobStatus(job)
	t.Log("Job is FAILED if a task is FAILED")
	assert.Nil(t, err)
	assert.Equal(t, job.Status, resources.Failed)

	// mark one task as success. should not mean success
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusSucceeded)
		break
	}
	err = jm.UpdateJobStatus(job)
	t.Log("One task SUCCESS does not result in job SUCCESS")
	assert.Nil(t, err)
	assert.NotEqual(t, job.Status, resources.Succeded)

	// mark all tasks as success. should mean job success
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusSucceeded)
	}
	err = jm.UpdateJobStatus(job)
	t.Log("Job is SUCCESSFUL if all tasks are SUCCESSFUL")
	assert.Nil(t, err)
	assert.Equal(t, job.Status, resources.Succeded)

	// mark one task as failed, others are successful. Still means failed
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusFailed)
		break
	}
	err = jm.UpdateJobStatus(job)
	t.Log("Job is FAILED if any task FAILS")
	assert.Nil(t, err)
	assert.Equal(t, job.Status, resources.Failed)
}

// TestCancelUpdates ensures that a cancelling a job works
// and that the following updates behave as expected
func TestCancelUpdates(t *testing.T) {
	batchClient := &mockBatchClient{
		map[string]resources.Task{},
	}
	jm := BatchJobManager{
		batchClient,
		store.NewMemoryStore(),
	}
	wf := resources.KitchenSinkWorkflow(t)
	input := []string{"test-start-input"}

	job, err := jm.CreateJob(wf, input)
	assert.Nil(t, err)

	// mark all tasks as running
	for _, task := range job.Tasks {
		task.SetStatus(resources.TaskStatusRunning)
	}
	err = jm.UpdateJobStatus(job)
	assert.NotEqual(t, resources.Cancelled, job.Status)

	// cancel the job
	t.Log("CancelJob marks a job as Cancelled")
	err = jm.CancelJob(job, "testing")
	assert.Nil(t, err)
	assert.Equal(t, resources.Cancelled, job.Status)

	// UpdateStatus ensures that job is still marked as Cancelled
	err = jm.UpdateJobStatus(job)
	assert.Nil(t, err)
	assert.Equal(t, resources.Cancelled, job.Status)
}

// TestCreateJob tests that tasks are created for a job in the right order
// with the appropriate settings
func TestCreateJob(t *testing.T) {
	mockClient := &mockBatchClient{
		map[string]resources.Task{},
	}
	jm := BatchJobManager{
		mockClient,
		store.NewMemoryStore(),
	}

	wf := resources.KitchenSinkWorkflow(t)
	input := []string{"test-start-input", "arg2"}

	job, err := jm.CreateJob(wf, input)
	assert.Nil(t, err)

	assert.Equal(t, len(job.Tasks), len(job.Workflow.States()))

	t.Log("Input data is passed to the first task only")
	assert.NotEmpty(t, job.Tasks[0].Input, mockClient.tasks[job.Tasks[0].ID].Input)
	assert.Empty(t, job.Tasks[1].Input, mockClient.tasks[job.Tasks[1].ID].Input)
	assert.Equal(t, job.Tasks[0].Input, mockClient.tasks[job.Tasks[0].ID].Input)
}