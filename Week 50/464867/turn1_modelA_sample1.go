package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Unbuffered I/O
	fmt.Print("Unbuffered Read:\n")
	for {
		byteData, err := file.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("%c", byteData)
	}
	fmt.Println()
}
