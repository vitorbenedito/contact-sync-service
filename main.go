package main

import (
	"contact-sync-service/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./api/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Contact Sync WebServer")
	log.Println("Initializing application")
	a := api.App{}
	a.Initialize()

	log.Println("Starting WebServer")
	a.Run(":8080")
}
