package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"hello-world-project-template-go/app"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to start temporal client", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete workrlfow", err)
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("getting result failed", err)

	}

	printResults(greeting, we.GetID(), we.GetRunID())

}

func printResults(greeting, workflowID, runID string) {

	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
