package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/superluminar-io/appsync-example-lambda/data"
)

type personPayload struct {
	ID int `json:"id"`
}

func handle(payload personPayload) (data.Person, error) {
	return data.PersonByID(payload.ID)
}

func main() {
	lambda.Start(handle)
}
