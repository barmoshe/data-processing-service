package main

import (
	"context"
	"data-processing-service/processing"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <data>")
	}
	inputData := os.Args[1]

	// Create a Temporal client using the new Dial method.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Generate a unique Workflow ID.
	workflowID := "data-processing-workflow-" + uuid.New().String()

	// Set up Workflow start options.
	workflowOptions := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "data-processing-task-queue",
	}

	// Start the Workflow.
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, processing.DataProcessingWorkflow, inputData)
	if err != nil {
		log.Fatalln("Unable to execute Workflow", err)
	}

	// Wait for the Workflow to complete and get the result.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	fmt.Println("Processed Data:", result)
}
