package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorldHandler is a simple HTTP handler.
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you could perform any business logic needed
	w.Write([]byte("Hello, World!"))
}

// CorsMiddleware sets CORS headers on HTTP responses.
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow all origins for demonstration purposes
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create the default router
	mux := http.NewServeMux()

	// Register a handler
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap the mux with the logging middleware
	loggedMux := CorsMiddleware(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
