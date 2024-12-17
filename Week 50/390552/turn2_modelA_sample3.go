package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorldHandler)

	// Use CORS handler from Gorilla
	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsMiddleware := handlers.CORS(corsObj)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", corsMiddleware(mux)); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
