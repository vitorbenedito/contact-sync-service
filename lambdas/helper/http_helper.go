package helper

import (
	"bytes"
	"contact-sync-service/domains"
	"encoding/json"

	"golang.org/x/exp/maps"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func BuildResponse(contacts map[string]domains.Contact) (Response, error) {
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
