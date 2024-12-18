package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

var fileCh chan struct{}

func init() {
	fileCh = make(chan struct{}, 1) // Buffer of size 1 to allow one goroutine at a time
}

func readFile(filename string) {
	select {
	case fileCh <- struct{}{}:
		defer func() { <-fileCh }()

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		fmt.Println("Read from file:", string(data))

	default:
		fmt.Println("File access already in progress.")
	}
}

func writeFile(filename string, content string) {
	select {
	case fileCh <- struct{}{}:
		defer func() { <-fileCh }()

		err := ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}

		fmt.Println("Written to file:", content)

	default:
		fmt.Println("File access already in progress.")
	}
}

func main() {
	filename := "concurrent_access.txt"

	// Create a few goroutines to read and write
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// Mix reads and writes randomly
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			if rand.Intn(2) == 0 { // Randomly decide to read or write
				readFile(filename)
			} else {
				writeFile(filename, fmt.Sprintf("Data %d\n", time.Now().UnixNano()))
			}
		}()
	}

	wg.Wait()
}
