package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func doWorkWithShutdown(id int, wg *sync.WaitGroup, ch chan<- string, errCh chan<- error, abort chan struct{}) {
	defer wg.Done()

	select {
	case <-abort:
		return // Exit early if an abort signal is received
	case <-time.After(time.Duration(id) * time.Second):
		if id%2 == 0 {
			ch <- fmt.Sprintf("Goroutine %d finished successfully", id)
		} else {
			errCh <- errors.New(fmt.Sprintf("Goroutine %d encountered an error", id))
		}
	}
}

func main() {
	var wg sync.WaitGroup
	resultsCh := make(chan string)
	errorsCh := make(chan error)
	abortCh := make(chan struct{})

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go doWorkWithShutdown(i, &wg, resultsCh, errorsCh, abortCh)
	}

	ok := true
	go func() {
		wg.Wait()
		close(resultsCh)
		close(errorsCh)
		ok = false // Set to false to indicate done
	}()

	for ok {
		select {
		case result, ok := <-resultsCh:
			if ok {
				fmt.Println("Result:", result)
			}
		case err, ok := <-errorsCh:
			if ok {
				fmt.Println("Error:", err)
				close(abortCh) // Signaling all goroutines to stop
			}
		}
	}
}
