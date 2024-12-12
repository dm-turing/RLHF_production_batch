package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, id int) {
	select {
	case <-ctx.Done():
		fmt.Println("Work cancelled for item", id)
		return
	default:
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Println("Completed work for item", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 5; i++ {
		go doWork(ctx, i)
	}

	time.Sleep(4 * time.Second) // Simulate more work than time allowed
}
