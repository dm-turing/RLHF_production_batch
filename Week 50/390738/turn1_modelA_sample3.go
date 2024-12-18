package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "file.txt"

	go writeData(fileName, "Goroutine 1\n")
	go writeData(fileName, "Goroutine 2\n")

	time.Sleep(1 * time.Second)

	readFile(fileName)
}

func writeData(fileName string, data string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func readFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(b))
}
