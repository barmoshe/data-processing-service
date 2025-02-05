package processing

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// DataProcessingWorkflow is the Workflow definition.
// It calls two Activities in sequence: first to add a random prefix (handled by Python),
// then to append a suffix (handled by Go).
func DataProcessingWorkflow(ctx workflow.Context, data string) (string, error) {
	// Define a retry policy for the Python activity.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        5, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	// Set ActivityOptions for the Python activity.
	pythonActivityOptions := workflow.ActivityOptions{
		TaskQueue:           "python-task-queue", // Points to the Python worker.
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}
	pythonCtx := workflow.WithActivityOptions(ctx, pythonActivityOptions)

	var prefixed string
	// Call the Python activity by name.
	err := workflow.ExecuteActivity(pythonCtx, "PythonAddRandomPrefixActivity", data).Get(pythonCtx, &prefixed)
	if err != nil {
		return "", fmt.Errorf("failed to add prefix: %w", err)
	}

	// Set ActivityOptions for the Go suffix activity.
	goActivityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	goCtx := workflow.WithActivityOptions(ctx, goActivityOptions)
	var uppercased string
	err = workflow.ExecuteActivity(goCtx, AddSuffixActivity, uppercased).Get(goCtx, &uppercased)
	if err != nil {
		return "", fmt.Errorf("failed to add suffix: %w", err)
	}
	// Set ActivityOptions for the ts UpperCase activity.
	tsActivityOptions := workflow.ActivityOptions{
		TaskQueue:           "typescript-task-queue", // TS worker's task queue.
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}
	tsCtx := workflow.WithActivityOptions(ctx, tsActivityOptions)
	var processed string

	err = workflow.ExecuteActivity(tsCtx, "TypeScriptToUppercaseActivity", prefixed).Get(tsCtx, &processed)
	if err != nil {
		return "", fmt.Errorf("failed to uppercase data: %w", err)
	}

	return processed, nil
}
