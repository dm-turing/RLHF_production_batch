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

// CORS Middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                              // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")            // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Custom-Header") // Allowed headers

		if r.Method == http.MethodOptions {
			// Preflight request
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create the default router
	mux := http.NewServeMux()

	// Register a handler
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap the mux with the logging middleware
	loggedMux := CORSMiddleware(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
