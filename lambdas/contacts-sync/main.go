package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/exp/maps"

	"contact-sync-service/services"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {

	log.Println("Received request from API Gateway - Sync Contacts API")

	contacts, _ := services.SyncContacts()

	var buf bytes.Buffer

	var response = make(map[string]any)
	response["syncedContacts"] = len(contacts)
	response["contacts"] = maps.Values(contacts)

	body, err := json.Marshal(response)

	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
