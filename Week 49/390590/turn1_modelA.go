package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Send result
	}
}

func main() {
	const numWorkers = 3
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to signal no more jobs

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel when done
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
