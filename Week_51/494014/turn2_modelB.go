package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func concurrentMergeMaps(maps ...map[string]int) map[string]int {
	// Create a channel to collect results from worker goroutines
	resultChan := make(chan map[string]int)

	// Calculate the number of workers based on the number of cores
	numWorkers := runtime.NumCPU()

	// Calculate the size of each worker's chunk
	chunkSize := len(maps[0])/numWorkers + 1

	// Spawn worker goroutines
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := min((i+1)*chunkSize, len(maps[0]))
		go func(start, end int) {
			defer wg.Done()
			workerMap := make(map[string]int)
			for _, m := range maps {
				for key, value := range m {
					if start <= len(key) && len(key) < end {
						workerMap[key] = value
					}
				}
			}
			resultChan <- workerMap
		}(start, end)
	}

	// Wait for all worker goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Merge the results from the worker goroutines
	mergedMap := make(map[string]int)
	for result := range resultChan {
		for key, value := range result {
			mergedMap[key] = value
		}
	}

	return mergedMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateLargeMap(sz int) map[string]int {
	m := make(map[string]int)
	for i := range sz {
		m[strconv.Itoa(i)] = i
	}
	return m
}

func main() {
	map1 := generateLargeMap(1000000)
	map2 := generateLargeMap(1000000)
	map3 := generateLargeMap(1000000)

	start := time.Now()
	concurrentMergeMaps(map1, map2, map3)
	fmt.Println("Concurrent Merge Time:", time.Since(start))

	start = time.Now()
	mergeMaps(map1, map2, map3)
	fmt.Println("Normal Merge Time:", time.Since(start))
}
