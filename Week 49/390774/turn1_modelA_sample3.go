package main

import (
	"errors"
	"fmt"
	"time"
)

func doWorkWithCancellation(id int, ch chan<- string, errCh chan<- error, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	time.Sleep(time.Duration(id) * time.Second)
	if id%2 == 0 {
		ch <- fmt.Sprintf("Goroutine %d finished successfully", id)
	} else {
		errCh <- errors.New(fmt.Sprintf("Goroutine %d encountered an error", id))
	}
}

func main() {
	resultsCh := make(chan string)
	errorsCh := make(chan error)
	done := make(chan struct{})

	for i := 1; i <= 5; i++ {
		go doWorkWithCancellation(i, resultsCh, errorsCh, done)
	}

	for {
		select {
		case result := <-resultsCh:
			fmt.Println("Result:", result)
		case err := <-errorsCh:
			fmt.Println("Error:", err)
			return // Stop processing on error
		case <-done:
			// All goroutines finished
			return
		}
	}
}
