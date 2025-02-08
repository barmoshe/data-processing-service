package main

import (
	"data-processing-service/activities"
	"data-processing-service/workflows"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create a Temporal client with the default options.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	taskQueue := "data-processing-task-queue"

	// Create a Worker that listens on the specified Task Queue.
	w := worker.New(c, taskQueue, worker.Options{})

	// Register the Workflow and the real Go suffix Activity with the Worker.
	w.RegisterWorkflow(workflows.DataProcessingWorkflow)
	w.RegisterActivity(activities.AddSuffixActivity)
	// Note: The Python and ts activities  will be handled by the Python/ts worker and is not registered here.

	// Start the Worker. This call blocks until the Worker is interrupted.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
