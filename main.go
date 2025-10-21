package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {

	// return "Hello, AWS Lambda!", nil
	return fmt.Sprintf("Hello, AWS Lambda!"), nil
}

func main() {
	lambda.Start(handler)
}
