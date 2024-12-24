package main

import (
	"log"
	"net/http"
)

// ErrorMiddleware is a middleware that handles errors returned by handlers.
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulate an error condition
		panic("Something went wrong!")
	})
	log.Println("Server starting on port :8000")
	log.Fatal(http.ListenAndServe(":8000", ErrorMiddleware(http.DefaultServeMux)))
}
