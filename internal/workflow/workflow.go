package workflow

import (
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	"github.com/jamillosantos/cadence-study/internal/workflow/activities"
)

type ReserveRequest struct {
	TransactionTraceID string
	OperationTraceID   string
	Amount             int64
	CurrencyCode       string
}

type ReserveResponse struct {
	ResponseCode  string `json:"response_code"`
	TransactionID string `json:"transaction_id"`
}

const (
	ReserveWorkflowName = "reserveWorkflow"
)

func ReserveWorkflow(ctx workflow.Context, request ReserveRequest) (ReserveResponse, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Reserve workflow started")

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var createDatabaseEntriesResponse activities.CreateDatabaseEntriesResponse
	err := workflow.ExecuteActivity(ctx, activities.CreateDatabaseEntriesActivity, activities.CreateDatabaseEntriesRequest{
		TransactionTraceID: request.TransactionTraceID,
		Amount:             request.Amount,
		CurrencyCode:       request.CurrencyCode,
	}).Get(ctx, &createDatabaseEntriesResponse)
	if err != nil {
		logger.Error("create transaction failed", zap.Error(err))
		return ReserveResponse{}, err
	}

	var doExternalRequestResponse activities.DoExternalRequestResponse

	err = workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Minute,
	}), activities.DoExternalRequestActivity, activities.DoExternalRequestRequest{
		Amount:       request.Amount,
		CurrencyCode: request.CurrencyCode,
	}).Get(ctx, &doExternalRequestResponse)
	if err != nil {
		logger.Error("doing external request failed", zap.Error(err))
		return ReserveResponse{}, err
	}

	logger.Info("Workflow completed.", zap.String("transaction_id", createDatabaseEntriesResponse.TransactionID))

	return ReserveResponse{
		TransactionID: createDatabaseEntriesResponse.TransactionID,
		ResponseCode:  doExternalRequestResponse.ResponseCode,
	}, nil
}
