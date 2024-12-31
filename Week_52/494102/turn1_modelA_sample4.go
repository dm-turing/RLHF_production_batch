package main

import (
	"context"
	"fmt"
	"time"
)

func exampleFunction(ctx context.Context) {
	fmt.Println("Function executing...")
	time.Sleep(10 * time.Second) // Simulate work
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	go exampleFunction(ctx)

	// Wait for the function to finish
	time.Sleep(10 * time.Second)
}
