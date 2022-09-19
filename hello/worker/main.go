package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"hello-world-project-template-go/app"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("unable to create temporal client", err)

	}
	defer c.Close()

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreetings)
	w.RegisterActivity(app.RunDocker)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}

}
