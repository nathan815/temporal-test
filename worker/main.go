package main

import (
	"log"

	"github.com/nathan815/temporal-hello-world/config"
	"github.com/nathan815/temporal-hello-world/workflows/getpost"
	"github.com/nathan815/temporal-hello-world/workflows/goroutinewf"
	"github.com/nathan815/temporal-hello-world/workflows/sleepingwf"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, config.MainTaskQueue, worker.Options{})
	w.RegisterWorkflow(sleepingwf.SleepingWorkflow)
	w.RegisterWorkflow(getpost.GetPostWithUser)
	w.RegisterWorkflow(goroutinewf.BasicGoroutineWorkflow)

	w.RegisterActivity(getpost.GetPost)
	w.RegisterActivity(getpost.GetUser)
	w.RegisterActivity(goroutinewf.Step1)
	w.RegisterActivity(goroutinewf.Step2)
	w.RegisterActivity(goroutinewf.Step3)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
