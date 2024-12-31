package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := process(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func process(ctx context.Context) error {
	go func() {
		time.Sleep(2 * time.Second)
		// cancel() // Simulate some other operation that cancels the context
	}()

	err := expensiveOperation(ctx)
	if err != nil {
		return err
	}

	return nil
}

func expensiveOperation(ctx context.Context) error {
	// Check for cancellation at regular intervals
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.New("cancelled")
		case <-ticker.C:
			fmt.Println("Doing work...")
		}
	}
}
