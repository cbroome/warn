package main

import (
	"context"
	"os"
	"strings"
	"warn/warn/states"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	state := os.Getenv("STATE")

	if strings.Compare(state, "maryland") == 0 {
		states.MarylandHandler();
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            state,
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}