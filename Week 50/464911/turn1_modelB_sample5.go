package main

import (
	"context"
	"fmt"
	"time"
)

func SomeOperation(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Perform the operation
		return nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := SomeOperation(ctx)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation completed successfully.")
	}
}
