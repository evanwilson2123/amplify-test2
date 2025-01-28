package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
  Name string `json:"name"`
}

type Response struct {
  Message string `json:"message"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
  return fmt.Sprintf("Hello %s!", name.Name ), nil
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  response := Response{
    Message: "Hello from Golang!",
  }

  body, err := json.Marshal(response)
  if err != nil {
    return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
  }

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusOK,
    Body: string(body),
  }, nil
}

func main() {
  lambda.Start(handler)
}
