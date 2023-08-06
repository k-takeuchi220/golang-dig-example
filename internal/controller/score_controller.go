package controller

import (
	"log"
	"sam-app/pkg/app/request"
	"sam-app/pkg/app/response"
	"sam-app/pkg/app/usecase"
	"sam-app/pkg/domain/model"

	"github.com/aws/aws-lambda-go/events"
)

type ScoreController struct {
	scoreUseCase usecase.ScoreUseCase
}

func NewScoreController(suc usecase.ScoreUseCase) *ScoreController {
	return &ScoreController{
		scoreUseCase: suc,
	}
}

func (c *ScoreController) CreateScore(r *request.CreateScoreRequest) (events.APIGatewayProxyResponse, error) {
	model := &model.Score{
		Name:  r.Name,
		Score: r.Score,
	}

	err := c.scoreUseCase.Create(model)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := response.ConvertToCreateScoreResponse()
	if err != nil {
		log.Fatal(err)
	}

	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Content-Type, Authorization",
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		Headers:    headers,
		StatusCode: 200,
	}, nil
}
