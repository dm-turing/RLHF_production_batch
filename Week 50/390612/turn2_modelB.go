package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/DataDog/dd-trace-go/tracer"
	"github.com/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var (
	books = map[string]Book{
		"1": Book{"1", "To Kill a Mockingbird", "Harper Lee"},
		"2": Book{"2", "1984", "George Orwell"},
	}
)

func getBook(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the request parameter
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Start a span for the getBook operation
	span, _ := tracer.StartSpanFromContext(r.Context(), "getBook")
	defer span.Finish()

	// Simulate a database query by adding a delay
	time.Sleep(50 * time.Millisecond)

	book, ok := books[bookID]
	if !ok {
		http.Error(w, fmt.Sprintf("Book with ID '%s' not found", bookID), http.StatusNotFound)
		return
	}

	// Set some custom tags on the span for later analysis
	span.SetTag("book.title", book.Title)
	span.SetTag("book.author", book.Author)

	json.NewEncoder(w).Encode(book)
}

func main() {
	// Initialize DataDog tracer
	tracer.Start(tracer.WithServiceName("book-service"))
	defer tracer.Stop()

	r := mux.NewRouter()
	r.HandleFunc("/books/{id}", getBook).Methods("GET")

	fmt.Println("Book service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
