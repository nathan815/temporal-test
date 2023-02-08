package sleepingwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SleepingWorkflow(ctx workflow.Context, duration time.Duration) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("The time is now: ", workflow.Now(ctx))
	logger.Info("About to sleep for ", duration)
	workflow.Sleep(ctx, duration)
	logger.Info("Wake up!")
	logger.Info("The time is now: ", workflow.Now(ctx))
	return nil
}
