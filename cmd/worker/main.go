package main

import (
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"

	"github.com/jamillosantos/cadence-study/internal/transport/cadence"
	"github.com/jamillosantos/cadence-study/internal/workflow"
	"github.com/jamillosantos/cadence-study/internal/workflow/activities"
)

func main() {
	var h cadence.SampleHelper
	h.SetupServiceConfig()

	registerWorkflowAndActivity(&h)
	startWorkers(&h)

	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *cadence.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
		FeatureFlags: client.FeatureFlags{
			WorkflowExecutionAlreadyCompletedErrorEnabled: true,
		},
	}
	h.StartWorkers(h.Config.DomainName, workflow.ApplicationName, workerOptions)
}

func registerWorkflowAndActivity(
	h *cadence.SampleHelper,
) {
	h.RegisterWorkflowWithAlias(workflow.ReserveWorkflow, workflow.ReserveWorkflowName)
	h.RegisterActivity(activities.CreateDatabaseEntriesActivity)
	h.RegisterActivity(activities.DoExternalRequestActivity)
}
