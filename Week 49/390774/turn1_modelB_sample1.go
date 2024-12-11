package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(time.Duration(id) * time.Second)
	if id%2 == 0 {
		ch <- fmt.Errorf("worker %d failed", id)
	} else {
		ch <- nil
	}
}

func main() {
	var wg sync.WaitGroup
	errorCh := make(chan error)

	// Start workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, errorCh, &wg)
	}

	// Select to handle errors and channel closure
	select {
	case err := <-errorCh:
		if err != nil {
			fmt.Println("Error received:", err)
		} else {
			fmt.Println("No error received")
		}
	case <-time.After(10 * time.Second):
		fmt.Println("Timeout occurred")
	}

	wg.Wait()
	close(errorCh) // Cleanup
}
