package main

import (
	"log"

	"data-processing-service/processing" // Import the processing package.

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create a Temporal client using the recommended Dial method.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Define the Task Queue name.
	taskQueue := "data-processing-task-queue"

	// Create a Worker that listens on the specified Task Queue.
	w := worker.New(c, taskQueue, worker.Options{})

	// Register the Workflow and the suffix Activity with the Worker.
	w.RegisterWorkflow(processing.DataProcessingWorkflow)
	w.RegisterActivity(processing.AddSuffixActivity)
	// Note: Do NOT register processing.AddRandomPrefixActivity in Go.
	// That Activity will be handled by the Python worker.

	// Start the Worker. This call blocks until the Worker is interrupted.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
