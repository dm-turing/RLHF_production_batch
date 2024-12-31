// Import required packages
package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Define a User struct with fields for name, age, and email
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Define a slice of User structs to store users data
var users []User

// Function to handle GET requests at /users endpoint
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Set content-type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Marshal the users slice into JSON data
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data as response
	w.Write(jsonData)
}

// Function to handle POST requests at /users endpoint to add new users
func addUser(w http.ResponseWriter, r *http.Request) {
	// Set content-type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Decode the request body into a User struct
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Append the new user to the users slice
	users = append(users, user)

	// Marshal the new user into JSON data
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data as response
	w.Write(jsonData)
}

// Function to handle GET requests at /users/{id} endpoint to get user by id
func getUserByID(w http.ResponseWriter, r *http.Request) {
	// Set content-type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Parse the user ID from the request URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user by ID from the users slice
	for _, user := range users {
		if user.ID == id {
			// Marshal the user into JSON data
			jsonData, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Write the JSON data as response
			w.Write(jsonData)
			return
		}
	}

	// User not found, return 404 Not Found
	http.NotFound(w, r)
}

func main() {
	// Initialize the users slice with some sample data
	users = []User{
		{ID: 1, Name: "Alice", Age: 25, Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Age: 30, Email: "bob@example.com"},
	}

	// Create a new router instance
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", addUser).Methods("POST")

