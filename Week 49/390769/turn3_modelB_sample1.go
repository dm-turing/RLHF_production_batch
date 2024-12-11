package main

import (
	"errors"
	"fmt"
)

// Example of error wrapping
func someOperation() error {
	return fmt.Errorf("some operation failed: %w", errors.New("internal error"))
}

func main() {
	err := someOperation()
	if err != nil {
		fmt.Println("Error:", err) // Outputs: Error: some operation failed: internal error

		// Unwrap the error
		wrappedErr := errors.Unwrap(err)
		if wrappedErr != nil {
			fmt.Println("Unwrapped Error:", wrappedErr) // Outputs: Unwrapped Error: internal error
		}
	}
}
