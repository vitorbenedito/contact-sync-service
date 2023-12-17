package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"contact-sync-service/services"

	"contact-sync-service/lambdas/helper"
)

func Handler(ctx context.Context) (helper.Response, error) {

	log.Println("Received request from API Gateway - Sync Contacts API")

	contacts, _ := services.SyncContactsParallel()

	resp, _ := helper.BuildResponse(contacts)

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
