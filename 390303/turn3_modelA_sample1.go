package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		mu.Lock()
		defer mu.Unlock()

		// Simulate work
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 completed work")

		// Simulate an error
		if false { // Replace with actual condition
			ctx = context.WithError(ctx, errors.New("condition failed in goroutine 1"))
		}
	}(ctx)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 2 canceled due to timeout")
			return
		default:
			err := concurrentFunc(&mu, ctx)
			if err != nil {
				ctx = context.WithError(ctx, err)
			}
		}
	}(ctx)

	err := ctx.Err()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func concurrentFunc(mu *sync.Mutex, ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine 2 completed work")

	// Placeholder for condition that might cause an error
	return nil
}
