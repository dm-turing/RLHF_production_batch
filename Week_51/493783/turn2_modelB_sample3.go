package main

import (
	"fmt"
	"os"
)

func userInputHandler(dataProcessor func(string)) {
	// Read user input
	input := os.Args[1]
	dataProcessor(input)
}

func processUserInput(input string) {
	secret := "SuperSecret"
	// Process input without access to secret
	fmt.Println("User Input:", input)
}

func main() {
	userInputHandler(processUserInput)
}
