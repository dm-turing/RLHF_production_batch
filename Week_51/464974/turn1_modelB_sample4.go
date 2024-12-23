package main

import (
	"os"
)

func main() {
	file, err := os.Open("example1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read or write to the file here
}
