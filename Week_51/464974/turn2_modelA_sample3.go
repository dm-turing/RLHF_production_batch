package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func writeToFile(data []byte, filename string) error {
	log.Printf("Writing to file: %s\n", filename)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v\n", err)
		return err
	}
	log.Printf("File written successfully: %s\n", filename)
	return nil
}

func main() {
	data := []byte("Hello, World!")
	filename := "example.txt"

	err := writeToFile(data, filename)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
	}
}
