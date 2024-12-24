package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close() // Ensure the file is closed

	// Read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return data, nil
}

func main() {
	filePath := "example.txt"
	data, err := readFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
