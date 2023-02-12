package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/nathan815/temporal-hello-world/config"
	"github.com/nathan815/temporal-hello-world/workflows/goroutinewf"
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
		ID:        fmt.Sprintf("goroutinewf/%v", uuid.New()),
		TaskQueue: config.MainTaskQueue,
	}

	// Start the Workflow
	if len(os.Args) < 2 {
		log.Fatalln("int arg parallelism not provided")
	}
	parallelism, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("int arg parallelism is invalid: %v", err)
	}
	we, err := c.ExecuteWorkflow(context.Background(), options, goroutinewf.ThreeStepGoroutineWorkflow, parallelism)
	if err != nil {
		log.Fatalln("unable to start workflow: ", err)
	}

	fmt.Println("Starting workflow...")
	fmt.Println("ID: ", we.GetID())
	fmt.Println("Run ID: ", we.GetRunID())
}
