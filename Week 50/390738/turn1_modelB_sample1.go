package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

func writeFile(filename string, content string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func main() {
	var wg sync.WaitGroup
	filename := "race.txt"
	content := "This is a test content."

	wg.Add(2)
	go writeFile(filename, content, &wg)
	go writeFile(filename, content, &wg)

	wg.Wait()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Content read from file:", string(data))
}
