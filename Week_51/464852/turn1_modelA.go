package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Set up a file for journaling
	file, err := os.Create("journal.log")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	// Use a buffered writer
	writer := bufio.NewWriter(file)

	// Calculate concurrent operations
	concurrency := runtime.GOMAXPROCS(0)

	// Measure a fixed number of entries
	const numEntries = 100000

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Start the timer
	waitSync := time.Now()

	// Each worker writes a share of the entries
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := i * numEntries / concurrency; j < (i+1)*numEntries/concurrency; j++ {
				// Write an entry to the journal
				mu.Lock()
				entry := fmt.Sprintf("%d\n", j)
				if _, err := writer.Write([]byte(entry)); err != nil {
					// log errors
				}
				mu.Unlock()
			}
			// Flush the buffered writer
			if err := writer.Flush(); err != nil {
				// log errors
			}
		}()
	}

	// Wait for all workers to finish
	wg.Wait()

	// Measure total time
	elapsedTime := time.Since(waitSync)

	// Calculate statistics
	averageTimePerOperation := elapsedTime / time.Duration(numEntries)
	throughput := numEntries / float64(elapsedTime.Seconds())

	// Print the results
	fmt.Printf("Elapsed time: %v\n", elapsedTime)
	fmt.Printf("Average time per operation: %v\n", averageTimePerOperation)
	fmt.Printf("Throughput: %.2f entries/s\n", throughput)
}
