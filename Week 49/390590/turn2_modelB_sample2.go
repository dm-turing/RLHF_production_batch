package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 5
	numTasks   = 20
)

var wg sync.WaitGroup

func worker(id int, tasks chan int) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	tasks := make(chan int)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks)
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
	fmt.Println("All tasks completed.")
}
