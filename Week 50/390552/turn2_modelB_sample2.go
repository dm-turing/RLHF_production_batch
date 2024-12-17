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

// ModifyRequestMiddleware adds a new header to the incoming request.
func ModifyRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add a custom header to the request
		r.Header.Add("X-Modified-Header", "true")

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create the default router
	mux := http.NewServeMux()

	// Register a handler
	mux.HandleFunc("/", HelloWorldHandler)

	// Wrap the mux with the logging middleware
	loggedMux := ModifyRequestMiddleware(mux)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
