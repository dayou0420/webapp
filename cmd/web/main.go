package main

import (
	"log"
	"net/http"
)

type applicatin struct{}

func main() {
	// set up an app config
	app := applicatin{}

	// get application routes
	mux := app.routes()

	// print put a message
	log.Println("Starting server on port 8080...")

	// start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
