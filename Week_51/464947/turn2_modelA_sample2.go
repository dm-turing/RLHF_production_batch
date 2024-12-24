package main

import (
	"fmt"
	"os"
	"sync"
)

func writeFileConcurrently(filePath string, data []byte, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	defer func() { ch <- struct{}{} }() // Signal completion

	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("failed to create file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("failed to write to file:", err)
		return
	}

	fmt.Println("Written:", string(data))
}

func main() {
	filePath := "example.txt"
	data := []byte("Hello, IoT!")
	wg := &sync.WaitGroup{}
	ch := make(chan struct{}, 5) // Limit concurrency to 5

	// Spawn multiple goroutines to write to the file concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeFileConcurrently(filePath, data, wg, ch)
	}

	// Allow all goroutines to proceed
	for i := 0; i < 10; i++ {
		<-ch // Wait for each goroutine to complete
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
