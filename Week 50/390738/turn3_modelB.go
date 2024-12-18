package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

const (
	maxReads  = 5               // Maximum number of read operations
	timeLimit = 2 * time.Second // Time limit for the process
)

var (
	readWriteQueue = make(chan struct{}, 1) // Buffer of size 1 to allow one read/write operation at a time
	wg             sync.WaitGroup
)

func readFile(filename string) {
	defer wg.Done()
	select {
	case readWriteQueue <- struct{}{}: // Acquire the read/write lock
		defer func() { <-readWriteQueue }()

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Read from file:", string(data))

	case <-time.After(timeLimit): // If the time limit is exceeded, exit the goroutine
		fmt.Println("Time limit exceeded for read.")
	}
}

func writeFile(filename string, content string) {
	select {
	case readWriteQueue <- struct{}{}: // Acquire the read/write lock
		defer func() { <-readWriteQueue }()

		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Written to file:", content)

	case <-time.After(timeLimit): // If the time limit is exceeded, exit the goroutine
		fmt.Println("Time limit exceeded for write.")
	}
}

func main() {
	filename := "round_robin.txt"
	initialContent := "Initial data\n"

	err := ioutil.WriteFile(filename, []byte(initialContent), 0644)
	if err != nil {
		fmt.Println("Error initializing file:", err)
		return
	}

	wg.Add(maxReads)
	// Start the read goroutines
	for i := 0; i < maxReads; i++ {
		go readFile(filename)
	}

	// Start the write goroutine with a time limit
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Random delay before write
		writeFile(filename, fmt.Sprintf("New data %d\n", time.Now().UnixNano()))
	}()

	// Wait for all read operations to complete
	wg.Wait()
	fmt.Println("All read operations completed.")
}
