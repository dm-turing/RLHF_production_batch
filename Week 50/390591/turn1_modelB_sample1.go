package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type User struct {
	Email string `json:"email"`
}

func main() {
	// Example JSON data
	jsonData := []byte(`{"email": "user@example.com"}`)
	var user User
	if err := json.Unmarshal(jsonData, &user); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	//Validate email format using regex
	if !validateEmail(user.Email) {
		fmt.Println("Invalid email format")
	} else {
		fmt.Println("Email format is valid")
	}
}

func validateEmail(email string) bool {
	// Validate email format using a simple regex
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}
