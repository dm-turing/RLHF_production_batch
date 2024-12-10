package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	filePath = "example.txt"
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func writeToFileBuffered(content string) {
	mutex.Lock()
	defer mutex.Unlock() // Ensure the critical section is locked

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // Create a buffered writer
	_, err = writer.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if err := writer.Flush(); err != nil { // Ensure data is written to the file
		fmt.Println("Error flushing buffer:", err)
		return
	}
	wg.Done() // Signal that the work is done
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeToFileBuffered(fmt.Sprintf("Line %d", i+1))
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All writes complete.")
}
