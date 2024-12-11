package main

import (
	"errors"
	"fmt"
	"os"
)

// Function that simulates a failure
func ReadFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err) // Wrapping the original error
	}
	return nil
}

// Function that chain errors
func ReadAndProcessFile(filename string) error {
	if err := ReadFile(filename); err != nil {
		return fmt.Errorf("ReadAndProcessFile error: %w", err)
	}
	// Simulate processing here...
	return nil
}

// Unwrapping the error to reveal the underlying cause
func main() {
	err := ReadAndProcessFile("nonexistent.txt")
	// Directly unwrapping the contextual error
	if unwrappedErr := errors.Unwrap(err); unwrappedErr != nil {
		fmt.Printf("The original error is: %v\n", unwrappedErr)
	}
}
