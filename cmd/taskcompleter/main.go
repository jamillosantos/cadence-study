package main

import (
	"context"
	"encoding/hex"
	"errors"
	"os"

	"github.com/jamillosantos/cadence-study/internal/transport/cadence"
	"github.com/jamillosantos/cadence-study/internal/workflow/activities"
)

func main() {
	if len(os.Args) < 3 {
		panic("invalid usage: taskcompleter <task token> <response code>")
	}

	taskToken, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic("invalid taskToken")
	}

	responseCode := os.Args[2]

	var responseErr error
	if responseCode == "error" {
		responseErr = errors.New("this is an error")
		responseCode = ""
	}

	var h cadence.SampleHelper
	h.SetupServiceConfig()

	client, err := h.Builder.BuildCadenceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = client.CompleteActivity(ctx, taskToken, activities.DoExternalRequestResponse{
		ResponseCode: responseCode,
	}, responseErr)
	if err != nil {
		panic(err)
	}
}
