package main

import (
	"fmt"
	"log"
)

func processData(data []int) error {
	log.Printf("Processing data: %v\n", data)
	defer log.Printf("Processed data: %v\n", data)

	// Perform some processing
	for _, value := range data {
		if value == 0 {
			return fmt.Errorf("cannot process zero value")
		}
	}
	return nil
}

func main() {
	data := []int{1, 2, 3, 0, 4}
	err := processData(data)
	if err != nil {
		log.Printf("Error processing data: %v\n", err)
	}
}
