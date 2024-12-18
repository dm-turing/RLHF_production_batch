package main

import (
	"fmt"
)

type logger func(string)

func processData(data []string, log logger) {
	for _, item := range data {
		result := processItem(item)
		log(result)
	}
}

func processItem(item string) string {
	return fmt.Sprintf("Processed item: %s", item)
}

func consoleLogger(message string) {
	fmt.Println(message)
}

func fileLogger(message string) {
	fmt.Printf("File Log: %s\n", message)
}

func main() {
	data := []string{"a", "b", "c"}
	processData(data, consoleLogger)
	fmt.Println()
	processData(data, fileLogger)
}
