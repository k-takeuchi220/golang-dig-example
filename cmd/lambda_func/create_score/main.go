package main

import (
	"fmt"
	"sam-app/di"
	"sam-app/internal/controller"
	"sam-app/pkg/app/request"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	params, err := request.CreateScoreParams(r)
	if err != nil {
		fmt.Println(err)
	}

	container := di.BuildContainer()
	var c *controller.ScoreController

	if err := container.Invoke(func(sc *controller.ScoreController) {
		c = sc
	}); err != nil {
		fmt.Println("Error resolving dependencies:", err)
	}

	response, err := c.CreateScore(params)
	if err != nil {
		fmt.Println(err)
	}

	return response, err
}

func main() {
	lambda.Start(handler)
}
