package main

import (
	"context"
	"fmt"
	"time"
)

func exampleFunction(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context canceled") // This will print, but the error is ignored
	default:
		fmt.Println("Function executing...")
		time.Sleep(10 * time.Second) // Simulate work
		fmt.Println("Function completed")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go exampleFunction(ctx)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Canceling context...")
		cancel() // Cancel the context after 2 seconds
	case <-ctx.Done():
		fmt.Println("Context canceled")
	}

	// Wait for the goroutine to finish
	time.Sleep(10 * time.Second)
}
