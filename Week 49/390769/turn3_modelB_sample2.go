package main

import (
	"fmt"
)

// Custom error type for chaining
type ChainedError struct {
	message string
	cause   error
}

func (e *ChainedError) Error() string {
	return fmt.Sprintf("%s: %v", e.message, e.cause)
}

func (e *ChainedError) Unwrap() error {
	return e.cause
}

func someOperation() error {
	return fmt.Errorf("%v", "because some operation failed")
}

// Example of error chaining
func anotherOperation() error {
	return &ChainedError{
		message: "another operation failed",
		cause:   someOperation(),
	}
}

func main() {
	err := anotherOperation()
	if err != nil {
		fmt.Println("Error:", err) // Outputs: Error: another operation failed: some operation failed: internal error

		// Unwrap the error chain
		for {
			if chainedErr, ok := err.(*ChainedError); ok {
				fmt.Println("Chained Error:", chainedErr.message)
				err = chainedErr.Unwrap()
			} else {
				break
			}
		}
		if err != nil {
			fmt.Println("Final Error:", err) // Outputs: Final Error: internal error
		}
	}
}
