package activities

import (
	"context"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

type DoExternalRequestRequest struct {
	Amount       int64
	CurrencyCode string
}

type DoExternalRequestResponse struct {
	ResponseCode string
}

func DoExternalRequestActivity(ctx context.Context, request DoExternalRequestRequest) (DoExternalRequestResponse, error) {
	logger := activity.GetLogger(ctx)
	ai := activity.GetInfo(ctx)

	logger.Info("doing external request started", zap.String("task_token", hex.EncodeToString(ai.TaskToken)))

	time.Sleep(time.Millisecond*100 + time.Duration(rand.Intn(1000)))

	if request.Amount%3 == 2 {
		return DoExternalRequestResponse{}, errors.New("failed doing external request")
	}

	return DoExternalRequestResponse{}, activity.ErrResultPending
}
