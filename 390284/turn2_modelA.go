package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	numWorkers = 10
	numTasks   = 10000
)

// Simulate a long-running task
func task(id int) {
	time.Sleep(time.Millisecond) // Simulates work
}

func main() {
	// Set GOMAXPROCS to utilize all available CPU cores
	fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())

	// Measure performance
	for _, procs := range []int{1, runtime.NumCPU()} {
		runtime.GOMAXPROCS(procs)

		start := time.Now()
		var wg sync.WaitGroup
		tasks := make(chan int, numTasks)

		// Start worker goroutines
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for taskID := range tasks {
					task(taskID)
				}
			}()
		}

		// Send tasks to the channel
		for i := 0; i < numTasks; i++ {
			tasks <- i
		}
		close(tasks)

		wg.Wait() // Wait for all workers to finish
		elapsed := time.Since(start)

		fmt.Printf("Tasks completed with GOMAXPROCS=%d in %s\n", procs, elapsed)
	}
}
