package main

import (
	"log"

	"github.com/pepeunlimited/go-twirp-starter-kit/internal/saga"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// A Worker is a service that executes Workflows and Activities.
// @see https://github.com/temporalio/documentation/blob/master/docs/go/workers.md
func main() {
	temporalClient, err := client.NewClient(client.Options{})
	if err != nil {
		log.Panicf("❌ temporal.client.NewClient occurred issue: %v", err)
	}
	defer temporalClient.Close()
	// worker hosts both Workflow and Activity functions
	temporalWorker := worker.New(temporalClient, saga.SagaTaskQueue, worker.Options{})
	activities := &saga.Activities{}
	temporalWorker.RegisterActivity(activities.Withdraw)
	temporalWorker.RegisterWorkflow(saga.SagaWorkflow)
	// Start listening to the Task Queue
	err = temporalWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Panicf("❌ temporal.worker.InterruptCh() occurred issue: %v", err)
	}
}
