package main

import (
	"fmt"
)

func validateHash(expected, computed string) {
	if expected != computed {
		panic("data tampered with; security breach detected")
	}
}

func main() {
	validateHash("expectedHash", "computedHash") // Will panic if hashes don't match
	fmt.Println("Data integrity verified.")
}
