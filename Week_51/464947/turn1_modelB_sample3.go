// Example 3: Avoiding Deferring Functions That Can Panic
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

	// Perform some operations that can panic
	doSomethingThatCanPanic()

	// The function will not return here if doSomethingThatCanPanic panics
	fmt.Println("Function completed successfully.")
}

func doSomethingThatCanPanic() {
	panic("Something unexpected happened!")
}
