package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	ch <- "Worker 1 completed"
}

func worker2(errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	errCh <- fmt.Errorf("Worker 2 failed")
}

func main() {
	var wg sync.WaitGroup
	resultCh := make(chan string)
	errorCh := make(chan error)

	// Start workers
	wg.Add(1)
	go worker1(resultCh, &wg)
	wg.Add(1)
	go worker2(errorCh, &wg)

	// Select to handle multiple channels
	select {
	case result := <-resultCh:
		fmt.Println("Result received:", result)
	case err := <-errorCh:
		fmt.Println("Error received:", err)
	default:
		fmt.Println("No messages received")
	}

	wg.Wait()
	close(resultCh)
	close(errorCh)
}
