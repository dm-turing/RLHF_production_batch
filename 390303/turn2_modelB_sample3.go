package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	errChan := make(chan error, 1)

	go func() {
		errChan <- processFile("example.txt")
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	// Perform file operations
	fmt.Println("Processing file...")

	// Simulate an error
	return errors.New("simulated error during file processing")
}
