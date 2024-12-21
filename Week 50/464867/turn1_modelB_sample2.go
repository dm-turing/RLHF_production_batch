package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter some text: ")
	var data [100]byte
	n, err := os.Stdin.Read(data[:])
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("You entered:", string(data[:n]))
}
