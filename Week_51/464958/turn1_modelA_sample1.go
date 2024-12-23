package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Ensure file is closed

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	data, err := readFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(data))
}
