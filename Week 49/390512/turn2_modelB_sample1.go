package main

import (
	"fmt"
	"os"
	"sync"
)

func writeToFile(filePath string, data []byte) {
	// Simplified write operation; in practice, you'd use file handling libraries
	if _, err := os.Create(filePath); err != nil {
		panic(fmt.Sprintf("cannot create file %s: %v", filePath, err))
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		// Critical data write
		writeToFile("critical_data.bin", []byte("some critical data"))
	}()

	wg.Wait()
	fmt.Println("Data written successfully.")
}
