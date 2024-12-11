package main

import (
	"fmt"
)

type MyError struct {
	Code      int
	Message   string
	Cause     error
	Timestamp string
	Severity  string
	Operation string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s: %v (severity: %s, operation: %s, timestamp: %s)",
		e.Code, e.Message, e.Cause, e.Severity, e.Operation, e.Timestamp)
}

// Example function that returns a custom error with metadata
func DoSomething() error {
	return &MyError{
		Code:      500,
		Message:   "Internal Server Error",
		Cause:     fmt.Errorf("failed to connect to database"),
		Timestamp: "2023-10-01T12:00:00Z",
		Severity:  "critical",
		Operation: "database connection",
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
