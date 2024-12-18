package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	fileName := "file.txt"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	var mutex sync.Mutex

	go writeData(&mutex, f, "Goroutine 1\n")
	go writeData(&mutex, f, "Goroutine 2\n")

	time.Sleep(1 * time.Second)

	f, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// scanner.Init(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func writeData(mutex *sync.Mutex, f *os.File, data string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, err := f.WriteString(data); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
