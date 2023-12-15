package main

import (
	"contact-sync-service/api"
	"log"
)

func main() {
	log.Println("Contact Sync WebServer")
	log.Println("Initializing application")
	a := api.App{}
	a.Initialize()

	log.Println("Starting WebServer")
	a.Run(":8080")
}
