package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter some text: ")
	var data [100]byte
	n, err := os.Stdin.Read(data[:])
	if err != nil && err != io.EOF {
		log.Fatalf("Error reading input: %v", err)
	}
	// Handle the EOF error separately if needed
	if err == io.EOF {
		fmt.Println("User input ended.")
	}
	fmt.Println("You entered:", string(data[:n]))
}
