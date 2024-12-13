package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d completed job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
