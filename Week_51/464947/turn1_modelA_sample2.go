package main

import (
	"fmt"
	"os"
)

func writeFile(filePath string, data []byte) error {
	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close() // Ensure the file is closed

	// Write data to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func main() {
	filePath := "example.txt"
	data := []byte("Hello, IoT!")
	err := writeFile(filePath, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File written successfully.")
}
