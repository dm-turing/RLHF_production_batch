package main

import (
	"fmt"
	"os"
)

func userInputHandler(callback func()) {
	// Read user input
	input := os.Args[1]
	callback()
}

func main() {
	secret := "SuperSecret"
	userInputHandler(func() {
		fmt.Println("User Input:", secret) // This is a vulnerability!
	})
}
