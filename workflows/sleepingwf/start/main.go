package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nathan815/temporal-hello-world/config"
	"github.com/nathan815/temporal-hello-world/workflows/sleepingwf"
	"go.temporal.io/sdk/client"
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("sleeping-wf-%v", uuid.New()),
		TaskQueue: config.MainTaskQueue,
	}

	// Start the Workflow
	duration := 12 * time.Hour
	we, err := c.ExecuteWorkflow(context.Background(), options, sleepingwf.SleepingWorkflow, duration)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	fmt.Println("Starting workflow...")
	fmt.Println("ID: ", we.GetID())
	fmt.Println("Run ID: ", we.GetRunID())
}
