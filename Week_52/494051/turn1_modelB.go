package main

import (
	"fmt"
	"sync"
	"time"
)

func generateReport(dataCh1 <-chan string, dataCh2 <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var report []string
	for {
		select {
		case data1, ok := <-dataCh1:
			if !ok {
				dataCh1 = nil // channel closed
			} else {
				report = append(report, data1)
			}
		case data2, ok := <-dataCh2:
			if !ok {
				dataCh2 = nil // channel closed
			} else {
				report = append(report, data2)
			}
		default:
			// If both channels are closed, break the loop
			if dataCh1 == nil && dataCh2 == nil {
				break
			}
		}
		// Simulate some work for the report generation
		time.Sleep(50 * time.Millisecond)
	}
	// Generate the final report and print it
	finalReport := fmt.Sprintf("Final Report:\n%s\n", report)
	fmt.Println(finalReport)
}

func main() {
	var wg sync.WaitGroup
	dataCh1 := make(chan string, 10)
	dataCh2 := make(chan string, 10)

	wg.Add(1)
	go generateReport(dataCh1, dataCh2, &wg)

	// Goroutines that produce data
	go func() {
		defer close(dataCh1)
		for i := 1; i <= 5; i++ {
			dataCh1 <- fmt.Sprintf("Data from source 1: %d", i)
		}
	}()

	go func() {
		defer close(dataCh2)
		for i := 1; i <= 3; i++ {
			dataCh2 <- fmt.Sprintf("Data from source 2: %d", i)
		}
	}()

	wg.Wait()
}
