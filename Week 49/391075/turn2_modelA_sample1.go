package main

import (
	"fmt"
	"sync"
	"time"
)

const numItems = 10_000

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	var wg sync.WaitGroup

	// Start the producer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numItems; i++ {
			ch <- i // Send to unbuffered channel
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

	fmt.Printf("Unbuffered channel elapsed time: %s\n", elapsed)
}
