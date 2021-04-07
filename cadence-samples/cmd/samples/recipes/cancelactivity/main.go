package main

import (
	"flag"
	"time"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"

	"github.com/uber-common/cadence-samples/cmd/samples/common"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func startWorkflow(h *common.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "cancel_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute * 30,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, sampleCancelWorkflow)
}

func cancelWorkflow(h *common.SampleHelper, wid string) {
	h.CancelWorkflow(wid)
}

func main() {
	var mode, wid string
	flag.StringVar(&mode, "m", "trigger", "Mode is worker, trigger or cancel.")
	flag.StringVar(&wid, "w", "<workflowID>", "w is the workflowID of the workflow to be canceled.")
	flag.Parse()

	var h common.SampleHelper
	h.SetupServiceConfig()

	switch mode {
	case "worker":
		h.RegisterWorkflow(sampleCancelWorkflow)
		h.RegisterActivity(activityToBeCanceled)
		h.RegisterActivity(activityToBeSkipped)
		h.RegisterActivity(cleanupActivity)
		startWorkers(&h)

		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	case "trigger":
		startWorkflow(&h)
	case "cancel":
		cancelWorkflow(&h, wid)
	}
}
