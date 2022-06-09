package activities

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

type CreateDatabaseEntriesRequest struct {
	TransactionTraceID string
	Amount             int64
	CurrencyCode       string
}

type CreateDatabaseEntriesResponse struct {
	TransactionID string
	OperationID   string
}

func CreateDatabaseEntriesActivity(ctx context.Context, request CreateDatabaseEntriesRequest) (CreateDatabaseEntriesResponse, error) {
	logger := activity.GetLogger(ctx)

	transactionID := uuid.NewV4().String()
	operationID := uuid.NewV4().String()
	logger.Info("CreateTransaction started", zap.String("transaction_trace_id", request.TransactionTraceID))

	return CreateDatabaseEntriesResponse{
		TransactionID: transactionID,
		OperationID:   operationID,
	}, nil
}
