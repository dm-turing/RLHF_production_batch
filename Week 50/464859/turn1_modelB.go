package main

import (
	"fmt"
	"log"
)

// Decorator function to handle errors
func HandleError(fn func() error) {
	if err := fn(); err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string) ([]byte, error) {
	// Simulate an error reading the file
	return nil, fmt.Errorf("error reading file %s", filename)
}

func main() {
	// Decorate the readFile function to handle errors
	HandleError(func() error {
		data, err := readFile("example.txt")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	})
}
