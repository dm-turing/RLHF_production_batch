package main

import (
	"fmt"
	"sync"
	"time"
)

func doWorkWithTimeout(id int, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	time.Sleep(time.Duration(id) * time.Second)
	if id%2 == 0 {
		fmt.Println("Goroutine", id, "finished successfully")
	} else {
		fmt.Println("Goroutine", id, "encountered an error")
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{}, 5)
	timeout := time.After(3 * time.Second)

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wg.Done()
			doWorkWithTimeout(id, done)
		}(i)
	}

	select {
	case <-timeout:
		fmt.Println("Operation timed out")
	case <-done:
		// At least one goroutine completed
	}

	wg.Wait()
	close(done)

	for range done {
		// Process completed goroutines
	}
}
