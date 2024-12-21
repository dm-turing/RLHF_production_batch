package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Fan-out: Create goroutines to work on the tasks
	const numWorkers = 5
	workChannel := make(chan int, numWorkers)
	resultChannel := make(chan string)

	// Start worker goroutines
	for w := 0; w < numWorkers; w++ {
		go worker(w, workChannel, resultChannel)
	}

	// Fan-in: Send tasks to the workers and collect results
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, task := range tasks {
		workChannel <- task
	}
	close(workChannel)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for result := range resultChannel {
		fmt.Println(result)
		wg.Done()
	}
	wg.Wait()
}

func worker(id int, workChannel <-chan int, resultChannel chan<- string) {
	for task := range workChannel {
		fmt.Printf("Worker %d received task %d\n", id, task)
		time.Sleep(time.Millisecond * 500) // Simulate work
		resultChannel <- fmt.Sprintf("Worker %d completed task %d", id, task)
	}
}
