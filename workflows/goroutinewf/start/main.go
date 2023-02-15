package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/nathan815/temporal-hello-world/config"
	"github.com/nathan815/temporal-hello-world/temporalext"
	"github.com/nathan815/temporal-hello-world/workflows/goroutinewf"
	"go.temporal.io/sdk/client"
)

func main() {
	c := temporalext.NewClient()
	defer c.Close()

	// Start the Workflow
	if len(os.Args) < 2 {
		log.Fatalln("int arg parallelism not provided")
	}
	parallelism, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("int arg parallelism is invalid: %v", err)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("goroutinewf/%v", uuid.New()),
		TaskQueue: config.MainTaskQueue,
	}
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, goroutinewf.ThreeStepGoroutineWorkflow, goroutinewf.ThreeStepGoroutineWorkflowInput{
		Parallelism: parallelism,
	})

	if err != nil {
		log.Fatalln("unable to start workflow: ", err)
	}

	fmt.Println("Starting workflow...")
	fmt.Println("ID: ", we.GetID())
	fmt.Println("Run ID: ", we.GetRunID())
}
