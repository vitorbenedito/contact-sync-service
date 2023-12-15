package main

import (
	"log"
)

func main() {
	log.Println("Contact Sync WebServer")
	log.Println("Initializing application")
	a := App{}
	a.Initialize()

	log.Println("Starting WebServer")
	a.Run(":8080")
}
