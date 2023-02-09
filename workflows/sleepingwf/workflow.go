package sleepingwf

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SleepingWorkflow(ctx workflow.Context, duration time.Duration) error {
	logger := workflow.GetLogger(ctx)
	logger.Info(fmt.Sprintf("The time is now: %v", workflow.Now(ctx)))
	logger.Info(fmt.Sprintf("About to sleep for %v", duration))
	workflow.Sleep(ctx, duration)
	logger.Info("Wake up!")
	logger.Info(fmt.Sprintf("The time is now: %v", workflow.Now(ctx)))
	return nil
}
