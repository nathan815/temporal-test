package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nathan815/temporal-hello-world/config"
	"github.com/nathan815/temporal-hello-world/workflows/getpost"
	"go.temporal.io/sdk/client"
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	postId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("invalid post id arg input. Error:", err)
	}

	options := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("get-post-%d", postId),
		TaskQueue: config.MainTaskQueue,
	}

	// Start the Workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, getpost.GetPostWithUser, postId)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	fmt.Println("Running workflow...")

	// Get the results
	var result getpost.PostWithUserOutput
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(result, we.GetID(), we.GetRunID())
}

func printResults(result getpost.PostWithUserOutput, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%+v", result)
}
