package goroutinewf

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

/**
* This sample workflow demonstrates how to use multiple Temporal gorotinues (instead of native goroutine) to process a
* a sequence of activities in parallel.
* In Temporal workflow, you should create goroutines using workflow.Go method.
 */

// SampleGoroutineWorkflow workflow definition
func BasicGoroutineWorkflow(ctx workflow.Context, parallelism int) (results []string, err error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 15 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	errors := make([]error, 0)

	for i := 0; i < parallelism; i++ {
		input1 := fmt.Sprint(i)

		workflow.Go(ctx, func(gCtx workflow.Context) {
			var result1 string
			err = workflow.ExecuteActivity(gCtx, Step1, input1).Get(gCtx, &result1)
			if err != nil {
				errors = append(errors, err)
				return
			}

			var result2 string
			err = workflow.ExecuteActivity(gCtx, Step2, result1).Get(gCtx, &result2)
			if err != nil {
				errors = append(errors, err)
				return
			}

			var result3 string
			err = workflow.ExecuteActivity(gCtx, Step3, result2).Get(gCtx, &result3)
			if err != nil {
				errors = append(errors, err)
				return
			}

			result := fmt.Sprintf("Result1: %v, Result2: %v, Result3: %v", result1, result2, result3)
			results = append(results, result)
		})
	}

	// Wait for Goroutines to complete. Await blocks until the condition function returns true.
	// The function is evaluated on every workflow state change. Consider using `workflow.AwaitWithTimeout` to
	// limit duration of the wait.
	_ = workflow.Await(ctx, func() bool {
		return err != nil || len(results) == parallelism
	})
	if len(errors) > 0 {
		return nil, fmt.Errorf("error occurred: %v", errors)
	}
	return
}

func Step1(input string) (output string, err error) {
	time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	return "Step1 Input: '" + input + "'", nil
}

func Step2(input string) (output string, err error) {
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	return "Step2 Input: (" + input + ")", nil
}

type Step3Heartbeat struct {
	Current int
	Total   int
	Time    time.Time
}

func Step3(ctx context.Context, input string) (output string, err error) {
	totalSleepSec := rand.Intn(20)
	for i := 0; i < totalSleepSec; i++ {
		time.Sleep(time.Second)
		activity.RecordHeartbeat(ctx, Step3Heartbeat{Current: i, Total: totalSleepSec, Time: time.Now()})
	}

	return "Step3 Input: (" + input + ")", nil
}
