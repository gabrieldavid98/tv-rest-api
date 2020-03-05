package main

import (
	"log"
	"os"
	"tv-rest-api/backend"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	connString := os.Getenv("CONNECTION_STR")

	httpBackend := backend.NewHTTPBackend(connString, 8080)
	httpBackend.Start()
}
