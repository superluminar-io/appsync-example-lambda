package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/superluminar-io/appsync-example-lambda/data"
)

func handle() (data.People, error) {
	return data.All(), nil
}

func main() {
	lambda.Start(handle)
}
