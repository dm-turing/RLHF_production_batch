package main

import (
	"errors"
	"fmt"
)

// Define a custom error type
type UserError struct {
	Cause  error
	Detail string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("UserError: %s: %v", e.Detail, e.Cause)
}

// Function that simulates user registration
func RegisterUser(username string) error {
	if username == "" {
		return &UserError{
			Cause:  errors.New("username cannot be empty"),
			Detail: "registration failed",
		}
	}
	// Simulate success
	return nil
}

// Function that simulates logging in
func LoginUser(username string) error {
	if username == "" {
		return &UserError{
			Cause:  errors.New("username cannot be empty"),
			Detail: "login failed",
		}
	}
	// Simulate success
	return nil
}

func main() {
	// Example of handling registration error
	if err := RegisterUser(""); err != nil {
		if userErr, ok := err.(*UserError); ok {
			fmt.Println("Handled UserError:", userErr)
		} else {
			fmt.Println("General error:", err)
		}
	}

	// Example of handling login error
	if err := LoginUser(""); err != nil {
		if userErr, ok := err.(*UserError); ok {
			fmt.Println("Handled UserError:", userErr)
		} else {
			fmt.Println("General error:", err)
		}
	}
}
