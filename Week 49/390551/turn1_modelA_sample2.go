package main

import (
	"fmt"
	"log"
	"os"
)

func readFile(filename string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Receiver error: %v", r)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("Failed to close file: %v", closeErr)
		}
	}()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return string(data), nil
}

func main() {
	data, err := readFile("non_existent_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
