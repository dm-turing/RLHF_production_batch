package main

import (
	"errors"
	"fmt"
	"time"
)

func doWork(id int, ch chan<- string, errCh chan<- error) {
	// Simulate work
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

	for i := 1; i <= 5; i++ {
		go doWork(i, resultsCh, errorsCh)
	}

	for i := 0; i < 5; i++ {
		select {
		case result := <-resultsCh:
			fmt.Println("Result:", result)
		case err := <-errorsCh:
			fmt.Println("Error:", err)
		}
	}
}
