package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func appendFile(filename string, content string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	filename := "truncation.txt"
	content := "Appended content."

	wg.Add(2)
	go appendFile(filename, content, &wg)
	go appendFile(filename, content, &wg)

	wg.Wait()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Content read from file:", string(data))
}
