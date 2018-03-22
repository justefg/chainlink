package services

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/adapters"
	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
)

// BeginRun creates a new run if the job is valid and starts the job.
func BeginRun(
	job models.JobSpec,
	initr models.Initiator,
	input models.RunResult,
	store *store.Store,
) (models.JobRun, error) {
	run, err := BuildRun(job, initr, store)
	if err != nil {
		return models.JobRun{}, err
	}
	return ExecuteRun(run, store, input)
}

// BuildRun checks to ensure the given job has not started or ended before
// creating a new run for the job.
func BuildRun(job models.JobSpec, i models.Initiator, store *store.Store) (models.JobRun, error) {
	now := store.Clock.Now()
	if !job.Started(now) {
		return models.JobRun{}, JobRunnerError{
			msg: fmt.Sprintf("Job runner: Job %v unstarted: %v before job's start time %v", job.ID, now, job.EndAt),
		}
	}
	if job.Ended(now) {
		return models.JobRun{}, JobRunnerError{
			msg: fmt.Sprintf("Job runner: Job %v ended: %v past job's end time %v", job.ID, now, job.EndAt),
		}
	}
	return job.NewRun(i), nil
}

// ExecuteRun starts the job and executes task runs within that job in the
// order defined in the run for as long as they do not return errors. Results
// are saved in the store (db).
func ExecuteRun(jr models.JobRun, store *store.Store, overrides models.RunResult) (models.JobRun, error) {
	jr.Status = models.RunStatusInProgress
	if err := store.Save(&jr); err != nil {
		return jr, wrapError(jr, err)
	}

	logger.Infow("Starting job", jr.ForLogger()...)
	unfinished := jr.UnfinishedTaskRuns()
	offset := len(jr.TaskRuns) - len(unfinished)
	latestRun := unfinished[0]

	merged, err := latestRun.Result.Merge(overrides)
	if err != nil {
		return jr, wrapError(jr, err)
	}
	latestRun.Result = merged

	for i, taskRunTemplate := range unfinished {
		taskRun, err := taskRunTemplate.MergeTaskParams(overrides.Data)
		if err != nil {
			return jr, wrapError(jr, err)
		}

		latestRun = markCompleted(startTask(taskRun, latestRun.Result, store)) // update to support blocking
		jr.TaskRuns[i+offset] = latestRun
		logTaskResult(latestRun, taskRun, i)

		if err := store.Save(&jr); err != nil {
			return jr, wrapError(jr, err)
		}
		if !latestRun.Status.Runnable() {
			break
		}
	}

	jr = jr.ApplyResult(latestRun.Result)
	logger.Infow("Finished current job run execution", jr.ForLogger()...)
	return jr, wrapError(jr, store.Save(&jr))
}

func logTaskResult(lr models.TaskRun, tr models.TaskRun, i int) {
	logger.Debugw("Produced task run", "taskRun", lr)
	logger.Debugw(fmt.Sprintf("Task %v %v", tr.Task.Type, tr.Result.Status), tr.ForLogger("task", i, "result", lr.Result)...)
}

func markCompleted(tr models.TaskRun) models.TaskRun {
	if tr.Status.Runnable() {
		return tr.MarkCompleted()
	}
	return tr
}

func startTask(
	tr models.TaskRun,
	input models.RunResult,
	store *store.Store,
) models.TaskRun {
	tr.Status = models.RunStatusInProgress

	adapter, err := adapters.For(tr.Task, store)
	if err != nil {
		tr.Status = models.RunStatusErrored
		tr.Result.SetError(err)
		return tr
	}

	return tr.ApplyResult(adapter.Perform(input, store))
}

func wrapError(run models.JobRun, err error) error {
	if err != nil {
		return fmt.Errorf("ExecuteRun: Job#%v: %v", run.JobID, err)
	}
	return nil
}

// JobRunnerError contains the field for the error message.
type JobRunnerError struct {
	msg string
}

// Error returns the error message for the run.
func (err JobRunnerError) Error() string {
	return err.msg
}
