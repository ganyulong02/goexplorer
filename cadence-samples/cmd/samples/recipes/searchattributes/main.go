package main

import (
	"context"
	"flag"
	"time"

	"go.uber.org/zap"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"

	"github.com/uber-common/cadence-samples/cmd/samples/common"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.SampleHelper) {
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}
	ctx := context.WithValue(context.Background(), CadenceClientKey, workflowClient)

	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope:              h.WorkerMetricScope,
		Logger:                    h.Logger,
		BackgroundActivityContext: ctx,
	}

	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func startWorkflow(h *common.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "searchAttributes_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
		SearchAttributes:                getSearchAttributesForStart(), // optional search attributes when start workflow
	}
	h.StartWorkflow(workflowOptions, searchAttributesWorkflow)
}

func getSearchAttributesForStart() map[string]interface{} {
	return map[string]interface{}{
		"CustomIntField": 1,
	}
}

func main() {
	var mode string
	flag.StringVar(&mode, "m", "trigger", "Mode is worker or trigger.")
	flag.Parse()

	var h common.SampleHelper
	h.SetupServiceConfig()

	switch mode {
	case "worker":
		h.RegisterWorkflow(searchAttributesWorkflow)
		h.RegisterActivity(listExecutions)
		startWorkers(&h)

		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	case "trigger":
		startWorkflow(&h)
	}
}
