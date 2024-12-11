package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, results []string, errors []error) {
	time.Sleep(time.Duration(id) * time.Second)
	if id%2 == 0 {
		results = append(results, fmt.Sprintf("Goroutine %d finished successfully", id))
	} else {
		errors = append(errors, errors.New(fmt.Sprintf("Goroutine %d encountered an error", id)))
	}
}

func main() {
	var wg sync.WaitGroup
	var results []string
	var errors []error

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(id int) {
			defer wg.Done()
			doWork(id, &results, &errors)
		}(i)
	}

	wg.Wait()

	for _, result := range results {
		fmt.Println("Result:", result)
	}
	for _, err := range errors {
		fmt.Println("Error:", err)
	}
}
