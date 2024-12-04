package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		mu.Lock()
		defer mu.Unlock()

		// Simulate work
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 completed work")

		// Simulate an error if required conditions aren't met
		if false { // Replace with actual condition
			errChan <- errors.New("condition failed in goroutine 1")
		}
	}()

	go func() {
		errChan <- concurrentFunc(&mu)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func concurrentFunc(mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()

	// Simulate work
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine 2 completed work")

	// Placeholder for condition that might cause an error
	return nil
}
