package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	errChan := make(chan error, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		defer close(errChan)
		errChan <- processFile(ctx)
	}(ctx)

	for err := range errChan {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func processFile(ctx context.Context) error {
	file, err := os.Open("example1.txt")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Perform file operations
		fmt.Println("Processing file...")

		// Simulate an error
		return errors.New("simulated error during file processing")
	}
}
