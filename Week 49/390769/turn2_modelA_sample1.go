package main

import (
	"fmt"
)

// Custom error type
type MyError struct {
	Code    int
	Message string
	Cause   error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s: %v", e.Code, e.Message, e.Cause)
}

// Example function that returns a custom error
func DoSomething() error {
	return &MyError{
		Code:    404,
		Message: "Resource not found",
		Cause:   fmt.Errorf("item with ID 123 does not exist"),
	}
}

func main() {
	err := DoSomething()
	if err != nil {
		if myErr, ok := err.(*MyError); ok {
			// Handle MyError specifically
			fmt.Printf("Custom Error: %s\n", myErr.Error())
			// You can also access the additional fields
			fmt.Printf("Error Code: %d\n", myErr.Code)
		} else {
			// Handle other types of errors
			fmt.Println("General error:", err)
		}
	}
}
