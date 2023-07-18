package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type WeatherResponse struct {
	Weather   string `json:"weather"`
	Timestamp int64  `json:"timestamp"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve dummy weather information
	weather := "Sunny"

	// Get current timestamp
	timestamp := time.Now().Unix()

	// Create the response object
	response := WeatherResponse{
		Weather:   weather,
		Timestamp: timestamp,
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, fmt.Errorf("failed to marshal response: %v", err)
	}

	// Return the response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseJSON),
	}, nil
}

func main() {
	lambda.Start(handler)
}
