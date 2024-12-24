package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func readFileConcurrently(filePath string, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close() // Ensure the file is closed

	// Read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	fmt.Println("Read:", string(data))
	return nil
}

func main() {
	filePath := "example.txt"
	wg := &sync.WaitGroup{}

	// Spawn multiple goroutines to read the file concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readFileConcurrently(filePath, wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
