package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func exampleFunction(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Return the context error
	default:
		fmt.Println("Function executing...")
		time.Sleep(10 * time.Second) // Simulate work
		return nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := exampleFunction(ctx); err != nil {
		log.Printf("Function failed: %v\n", err)
	}

	// Wait for the function to finish
	time.Sleep(10 * time.Second)
}
