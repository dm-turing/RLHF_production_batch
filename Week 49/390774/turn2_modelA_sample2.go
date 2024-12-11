package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func doWorkWithContext(ctx context.Context, id int, resultsCh chan<- string, errorsCh chan<- error) {
	time.Sleep(time.Duration(id) * time.Second)
	select {
	case <-ctx.Done():
		return // Return if the context has expired or been canceled
	default:
		if id%2 == 0 {
			resultsCh <- fmt.Sprintf("Goroutine %d finished successfully", id)
		} else {
			errorsCh <- errors.New(fmt.Sprintf("Goroutine %d encountered an error", id))
		}
	}
}

func main() {
	resultsCh := make(chan string)
	errorsCh := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 5; i++ {
		go doWorkWithContext(ctx, i, resultsCh, errorsCh)
	}

	for {
		select {
		case result := <-resultsCh:
			fmt.Println("Result:", result)
		case err := <-errorsCh:
			fmt.Println("Error:", err)
		case <-ctx.Done():
			fmt.Println("Context canceled or timed out!")
			return
		}
	}
}
