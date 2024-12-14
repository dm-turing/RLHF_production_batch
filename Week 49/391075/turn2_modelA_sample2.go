package main

import (
	"fmt"
	"sync"
	"time"
)

const numItems = 10_000
const bufferSize = 100 // Define buffer size

func main() {
	// Create a buffered channel
	ch := make(chan int, bufferSize)

	var wg sync.WaitGroup

	// Start the producer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numItems; i++ {
			ch <- i // Send to buffered channel
		}
		close(ch)
	}()

	// Start the consumer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// Receiving value
		}
	}()

	start := time.Now()
	wg.Wait() // Wait for both goroutines to finish
	elapsed := time.Since(start)

	fmt.Printf("Buffered channel elapsed time: %s\n", elapsed)
}
