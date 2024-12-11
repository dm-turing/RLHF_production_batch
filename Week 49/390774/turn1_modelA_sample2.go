package main

import (
	"errors"
	"fmt"
	"time"
)

func doWorkWithTimeout(id int, ch chan<- string, errCh chan<- error) {
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
	timeoutCh := time.After(3 * time.Second)

	for i := 1; i <= 5; i++ {
		go doWorkWithTimeout(i, resultsCh, errorsCh)
	}

	for i := 0; i < 5; i++ {
		select {
		case result := <-resultsCh:
			fmt.Println("Result:", result)
		case err := <-errorsCh:
			fmt.Println("Error:", err)
		case <-timeoutCh:
			fmt.Println("Operation timed out")
			return // or handle as desired
		}
	}
}
