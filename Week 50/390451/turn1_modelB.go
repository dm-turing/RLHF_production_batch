package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Access an environment variable
	cloudProvider := os.Getenv("CLOUD_PROVIDER")
	fmt.Println("Running on:", cloudProvider)
	http.HandleFunc("/health", healthCheckHandler)
	// Example of a health check endpoint
	// if err :=  err != nil {
	// log.Fatal("Error setting up health check handler:", err)
	// }

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Healthy")
}
