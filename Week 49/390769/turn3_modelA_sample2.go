package main

import (
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

func main() {
	err := ReadAndProcessFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)

		// Check for the original error
		if os.IsNotExist(err) {
			fmt.Println("The file does not exist")
		} else {
			fmt.Println("Some other error occurred")
		}
	}
}
