package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent a person
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Email:   "john.doe@example.com",
		IsAdmin: true,
	}

	// Marshal the Person struct into a JSON byte array
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println("JSON Data:", string(jsonData))
}
