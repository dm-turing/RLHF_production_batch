package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

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
		wrappedErr := errors.Wrap(err, "failed to read file")
		fmt.Printf("Wrapped error: %+v\n", wrappedErr)
	}
}
