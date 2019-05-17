package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/superluminar-io/appsync-example-lambda/data"
)

type personFriendsPayload struct {
	ID int `json:"id"`
}

func handle(payload personFriendsPayload) (data.People, error) {
	return data.FriendsByID(payload.ID)
}

func main() {
	lambda.Start(handle)
}
