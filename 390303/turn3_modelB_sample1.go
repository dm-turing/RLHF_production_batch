package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var mu sync.Mutex
	errChan := make(chan error, 2)

	go func() {
		errChan <- doWork(ctx, &mu, 2*time.Second, "Goroutine 1")
	}()

	go func() {
		errChan <- doWork(ctx, &mu, 1*time.Second, "Goroutine 2")
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func doWork(ctx context.Context, mu *sync.Mutex, workDuration time.Duration, name string) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("%s: context canceled: %w", name, ctx.Err())
	default:
		mu.Lock()
		defer mu.Unlock()
		// Simulate work
		time.Sleep(workDuration)
		fmt.Printf("%s completed work\n", name)
		return nil
	}
}
