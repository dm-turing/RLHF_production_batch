package main

import (
	"fmt"
	"io"
	"sync"
)

func main() {
	// Create a pipe
	pr, pw := io.Pipe()
	defer pw.Close() // Close the writer end of the pipe after use

	var wg sync.WaitGroup
	// Reader goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Read data from the pipe
		data, err := io.ReadAll(pr)
		if err != nil {
			fmt.Println("Error reading from pipe:", err)
			return
		}
		fmt.Println("Received data:", string(data))
	}()

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pw.CloseWithError(nil) // Close the writer end of the pipe with a non-nil error to signal the end of the stream
		// Write data to the pipe
		dataToWrite := "Hello, World!"
		n, err := pw.Write([]byte(dataToWrite))
		if err != nil {
			fmt.Println("Error writing to pipe:", err)
			return
		}
		fmt.Println("Written", n, "bytes to the pipe")
	}()
	wg.Wait() // Wait for both goroutines to finish
}
