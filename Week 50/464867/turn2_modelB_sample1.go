package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter some text: ")
	// Using io.ReadCloser to allow for network connections in the future
	reader := bufio.NewReaderSize(os.Stdin, 5)

	data, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			fmt.Println("User input ended.")
		} else {
			log.Fatalf("Error reading input: %v", err)
		}
	}
	fmt.Println("You entered:", data)
}
