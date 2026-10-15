package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Message string `json:"message"`
}

func hello() (response, error) {
	return response{Message: "Hello, world!"}, nil
}

func main() {
	lambda.Start(hello)
}
