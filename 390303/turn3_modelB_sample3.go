package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		errChan <- processFile(ctx, "example.txt")
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func processFile(ctx context.Context, filename string) error {
	file, err := os.Open(filename)
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
		// Simulate some file processing
		time.Sleep(1 * time.Second)
		fmt.Println("Processing file completed")
	}

	// Simulate an error
	return errors.New("simulated error during file processing")
}
