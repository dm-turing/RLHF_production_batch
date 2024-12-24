package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const numConcurrentOperations = 10

var (
	filePool  chan *os.File
	poolMutex sync.Mutex
)

func init() {
	filePool = make(chan *os.File, numConcurrentOperations)
	for i := 0; i < numConcurrentOperations; i++ {
		file, err := os.Open("example.txt")
		if err != nil {
			panic(err)
		}
		filePool <- file
	}
}

func readFileConcurrently() {
	// Acquire a file handle from the pool
	poolMutex.Lock()
	file := <-filePool
	poolMutex.Unlock()
	defer func() {
		// Return the file handle to the pool
		poolMutex.Lock()
		filePool <- file
		poolMutex.Unlock()
	}()

	// Read the file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	// Spawn concurrent goroutines to read the file
	for i := 0; i < numConcurrentOperations; i++ {
		go readFileConcurrently()
	}
}
