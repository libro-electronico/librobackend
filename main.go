package main

import (
	"libro-electronico/helper"
	"log"
	"net/http"
)

func main() {
    // Setup routes
    http.HandleFunc("/webhook", helper.WebHookHandler)

    // Menjalankan server
    log.Printf("Server starting on port %s")
    log.Fatal(http.ListenAndServe(":", nil))
}