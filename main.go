package main

import (
	"context"
	route "libro-electronico/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Mendapatkan port dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika tidak ada environment variable
	}

	// Inisialisasi router atau handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", route.URL) // Menggunakan handler dari package route

	// Membuat server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Channel untuk menangkap signal sistem
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Goroutine untuk menjalankan server
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v\n", port, err)
		}
	}()

	// Menunggu sinyal interrupt
	<-stop

	// Shutdown server dengan waktu tunggu untuk penyelesaian request
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
