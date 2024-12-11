package main

import (
	"errors"
	"fmt"
	"os"
)

// Custom error type
type MyError struct {
	Original error
	Message  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("MyError: %s: %v", e.Message, e.Original)
}

// Function that simulates a failure
func ReadFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err) // Wrapping the original error
	}
	return nil
}

func main() {
	err := ReadFile("nonexistent.txt")
	if err != nil {
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Println("Custom error:", myErr)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
