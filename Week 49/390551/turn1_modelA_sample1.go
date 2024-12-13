package main

import (
	"fmt"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func readFile(filename string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		handleError(file.Close())
	}()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	data, err := readFile("non_existent_file.txt")
	handleError(err)
	fmt.Println(data)
}
