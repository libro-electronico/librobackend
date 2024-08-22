package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	
	log.Printf("Server starting on port %s", port)
	
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
