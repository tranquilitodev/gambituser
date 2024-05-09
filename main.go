package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLmabda)
}
func EjecutoLmabda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation)(events.CognitoEventUserPoolsPostConfirmation,error) {

}
