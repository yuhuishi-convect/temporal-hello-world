package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	var result string

	err := workflow.ExecuteActivity(ctx, ComposeGreetings, name).Get(ctx, &result)
	if err != nil {
		return result, err
	}

	err = workflow.ExecuteActivity(ctx, RunDocker).Get(ctx, nil)
	return result, err

}
