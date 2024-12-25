package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Define a struct to represent the application's model
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

// Define a controller struct to handle HTTP requests and responses
type UserController struct {
}

// Implement the ServeHTTP method to handle GET requests for the "/user" endpoint
func (uc *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse the URL query parameters
	query := r.URL.Query()

	// Extract the query parameters and create a User struct
	user := User{
		Name:     query.Get("name"),
		Age:      parseInt(query.Get("age"), 0),
		Location: query.Get("location"),
	}

	// Validate the user data (optional)
	// For example, you can check if the required fields are present and have valid values.

	// Render the response using the view
	renderUser(w, user)
}

// Helper function to parse a string to an int, returning a default value if parsing fails
func parseInt(s string, defaultValue int) int {
	if value, err := strconv.Atoi(s); err == nil {
		return value
	}
	return defaultValue
}

// Define a view function to render the user data as JSON
func renderUser(w http.ResponseWriter, user User) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Create a new UserController instance
	uc := &UserController{}

	// Register the "/user" endpoint with the UserController
	http.Handle("/user", uc)

	// Start the HTTP server on port 8080
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
