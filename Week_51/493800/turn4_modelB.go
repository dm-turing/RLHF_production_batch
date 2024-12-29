package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Create a new file and write some data to it
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data := []byte("Hello, world!")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read the data from the file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	readData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Data read from file:", string(readData))
}
