// Example 1: Opening a File with Defer
package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file "example.txt" for reading
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Defer the file closure to ensure it's closed properly
	defer file.Close()

	// Read the file content and print it to the console
	// (Your file reading logic here)
}
