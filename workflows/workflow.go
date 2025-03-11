package workflows

import (
	"data-processing-service/activities"
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func DataProcessingWorkflow(ctx workflow.Context, data string) (string, error) {
	// Define a retry policy.
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,       // First retry after 1 second
		BackoffCoefficient: 2.0,               // Double the wait time on each retry (1s → 2s → 4s → 8s, etc.)
		MaximumInterval:    100 * time.Second, // Cap wait time at 100 seconds
		MaximumAttempts:    50,                // Retry up to 5 times before giving up
	}

	// Step 1: Add a prefix(Python)
	pythonCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "python-task-queue",
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	})

	var prefixed string
	if err := workflow.ExecuteActivity(pythonCtx, "PythonAddRandomPrefixActivity", data).Get(pythonCtx, &prefixed); err != nil {
		return "", fmt.Errorf("failed to add prefix: %w", err)
	}

	// Step 2: Add a suffix(Go)
	goCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	})

	var uppercased string
	if err := workflow.ExecuteActivity(goCtx, activities.AddSuffixActivity, prefixed).Get(goCtx, &uppercased); err != nil {
		return "", fmt.Errorf("failed to add suffix: %w", err)
	}

	// Step 3: upercase the string (ts)
	tsCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskQueue:           "typescript-task-queue",
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	})

	var processed string
	if err := workflow.ExecuteActivity(tsCtx, "TypeScriptToUppercaseActivity", uppercased).Get(tsCtx, &processed); err != nil {
		return "", fmt.Errorf("failed to uppercase data: %w", err)
	}

	return processed, nil
}
