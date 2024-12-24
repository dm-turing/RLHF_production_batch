package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const numConcurrentOperations = 10

var wg sync.WaitGroup

func readFile(filePath string) {
	defer wg.Done()

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	filePath := "example.txt"

	// Semaphore to limit concurrent file operations
	sem := make(chan struct{}, numConcurrentOperations)

	// Spawn concurrent goroutines to read the file
	for i := 0; i < numConcurrentOperations; i++ {
		wg.Add(1)
		go func() {
			sem <- struct{}{} // Acquire a semaphore
			defer func() {
				<-sem // Release the semaphore
			}()
			readFile(filePath)
		}()
	}

	wg.Wait() // Wait for all goroutines to complete
}
