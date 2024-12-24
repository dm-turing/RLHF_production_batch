// Example 2: Opening Multiple Files with Defer
package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file "example1.txt" for reading
	file1, err := os.Open("example1.txt")
	if err != nil {
		fmt.Println("Error opening file1:", err)
		return
	}

	// Defer the file1 closure to ensure it's closed properly
	defer file1.Close()

	// Open the file "example2.txt" for reading
	file2, err := os.Open("example2.txt")
	if err != nil {
		fmt.Println("Error opening file2:", err)
		return
	}

	// Defer the file2 closure to ensure it's closed properly
	defer file2.Close()

	// Read the file content from both files and print it to the console
	// (Your file reading logic here)
}
