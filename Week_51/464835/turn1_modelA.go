package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", errorHandler(http.DefaultServeMux)))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an error
	err := someDatabaseOperation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Hello, World!")
}

func someDatabaseOperation() error {
	// Simulate a database error
	return fmt.Errorf("database operation failed")
}

func errorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, err.(error).Error(), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
