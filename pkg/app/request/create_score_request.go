package request

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type CreateScoreRequest struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
}

func CreateScoreParams(r events.APIGatewayProxyRequest) (*CreateScoreRequest, error) {
	var req *CreateScoreRequest
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		return req, err
	}
	return req, nil
}
