package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Buffered I/O
	scanner := bufio.NewScanner(file)
	fmt.Print("Buffered Read:\n")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}
