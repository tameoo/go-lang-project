package main

import (
	"log"
	"net/http"
	router "rest-api-mid/infrastructure/router"
)

func main() {
	// creating server
	server := router.CreateRouter()

	// listening server on port localhost:8080
	log.Fatal(http.ListenAndServe(":8080", server))
}
