package main

import (
	"fmt"
	"sync"
)

type Task struct {
	id int
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.id)
		// Simulate work
		// time.Sleep(time.Second)
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	tasks := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Create tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{id: i}
	}
	close(tasks) // No more tasks will be sent

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks completed.")
}
