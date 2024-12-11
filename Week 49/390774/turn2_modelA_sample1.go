package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func doWork(id int, wg *sync.WaitGroup, ch chan<- string, errCh chan<- error) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * time.Second)
	if id%2 == 0 {
		ch <- fmt.Sprintf("Goroutine %d finished successfully", id)
	} else {
		errCh <- errors.New(fmt.Sprintf("Goroutine %d encountered an error", id))
	}
}

func main() {
	var wg sync.WaitGroup
	resultsCh := make(chan string)
	errorsCh := make(chan error)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go doWork(i, &wg, resultsCh, errorsCh)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(resultsCh)
		close(errorsCh)
	}()

	for {
		select {
		case result, ok := <-resultsCh:
			if ok {
				fmt.Println("Result:", result)
			}
		case err, ok := <-errorsCh:
			if ok {
				fmt.Println("Error:", err)
			}
		default:
			if len(resultsCh) == 0 && len(errorsCh) == 0 {
				return // Exit if all channels are closed
			}
		}
	}
}
