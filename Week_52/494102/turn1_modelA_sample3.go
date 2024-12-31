package main

import (
	"context"
	"fmt"
	"time"
)

func exampleFunction(ctx context.Context) {
	// Creating a new context for each operation
	childCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	fmt.Println("Function executing...")
	time.Sleep(1 * time.Second) // Simulate work

	// Another new context
	anotherChildCtx, anotherCancel := context.WithTimeout(ctx, 2*time.Second)
	defer anotherCancel()

	fmt.Println("Another function executing...")
	time.Sleep(1 * time.Second) // Simulate work
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go exampleFunction(ctx)

	// Wait for the function to finish
	time.Sleep(10 * time.Second)
}
