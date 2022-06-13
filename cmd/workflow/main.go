package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/satori/go.uuid"
	"go.uber.org/cadence/client"

	"github.com/jamillosantos/cadence-study/internal/transport/cadence"
	"github.com/jamillosantos/cadence-study/internal/workflow"
)

func main() {
	var h cadence.SampleHelper
	h.SetupServiceConfig()

	startWorkflow(&h)
}

func startWorkflow(h *cadence.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "helloworld_" + uuid.NewV4().String(),
		TaskList:                        workflow.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	transactionTraceID := uuid.NewV4().String()
	operationTraceID := uuid.NewV4().String()

	exec := h.StartWorkflow(workflowOptions, workflow.ReserveWorkflowName, workflow.ReserveRequest{
		TransactionTraceID: transactionTraceID,
		OperationTraceID:   operationTraceID,
		Amount:             123,
		CurrencyCode:       "GBP",
	})

	ctx := context.Background()
	var result workflow.ReserveResponse
	if err := h.GetWorkflow(ctx, exec.ID, &result); err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Print("Result: ")
	_ = json.NewEncoder(os.Stdout).Encode(&result)
	fmt.Println()
}
