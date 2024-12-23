package main

import (
	"fmt"
	"log"
)

func logFunctionExecution(name string) {
	log.Printf("Entering function: %s\n", name)
	defer log.Printf("Exiting function: %s\n", name)
	// The rest of the function logic goes here
}

func processData(data []byte) string {
	logFunctionExecution("processData")
	return string(data)
}

func main() {
	data := []byte("Hello, World!")
	result := processData(data)
	fmt.Println("Result:", result)
}
