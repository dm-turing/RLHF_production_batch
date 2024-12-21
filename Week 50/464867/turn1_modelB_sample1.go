package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a buffered reader with a buffer size of 5 bytes
	reader := bufio.NewReaderSize(os.Stdin, 5)
	fmt.Print("Enter some text: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("You entered:", data)
}
