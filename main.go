package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/liberoelectronico", func(w http.ResponseWriter, r *http.Request) {
        // Your handler logic here
        w.Write([]byte("Hello, World!"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}