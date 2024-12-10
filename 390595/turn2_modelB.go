package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/ulisesalvarado/flock"
)

var (
	filePath = pflag.StringP("file", "f", "example.txt", "path to the file")
)

func main() {
	pflag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Acquire an exclusive lock
	if err := flock.Lock(file); err != nil {
		fmt.Println("Error acquiring lock:", err)
		return
	}
	defer flock.Unlock(file) // Ensure the lock is released

	// Use buffered I/O
	writer := bufio.NewWriter(file)

	// Concurrently write to the file
	for i := 1; i <= 5; i++ {
		go func(i int) {
			_, err := writer.WriteString(fmt.Sprintf("Line %d\n", i))
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
		}(i)
	}

	// Flush the buffer to ensure all data is written to the file
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing buffer:", err)
	}

	fmt.Println("Data written to file successfully.")
}
