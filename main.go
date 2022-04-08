package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type RequestModel struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, requestModel RequestModel) (string, error) {
	if len(requestModel.Name) > 0 {
		return fmt.Sprintf("Hello %s!", requestModel.Name), nil
	}
	return "Hello world", nil
}

func main() {
	lambda.Start(HandleRequest)
}
