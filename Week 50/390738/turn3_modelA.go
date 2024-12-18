package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var fileCh chan struct{}
var writesComplete atomic.Bool
var lastWrite int64
var wg sync.WaitGroup

func init() {
	fileCh = make(chan struct{}, 1)
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
		atomic.StoreInt64(&lastWrite, time.Now().UnixNano())

	default:
		fmt.Println("File access already in progress.")
	}
}

func processUntilGracefulExit(filename string, totalWrites int, timeout time.Duration) {
	for !writesComplete.Load() {
		time.Sleep(5 * time.Millisecond)
		now := time.Now().UnixNano()

		if lastWrite > 0 && (now-lastWrite) > timeout.Nanoseconds() {
			writesComplete.Store(true)
			// atomic.StoreBool(&writesComplete, true)
			fmt.Println("No writes for a while, finishing reads.")
		}
	}

	fmt.Println("Reads in progress...")
	// Keep reading until all reads are done
	for {
		select {
		case fileCh <- struct{}{}:
			defer func() { <-fileCh }()

			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println("Error reading file during graceful exit:", err)
				continue
			}

			fmt.Println("Read from file during graceful exit:", string(data))

		default:
			time.Sleep(10 * time.Millisecond) // Wait for file to be available
		}
	}
}

func main() {
	filename := "round_robin_file.txt"
	totalWrites := 10 // Example number of writes
	timeout := time.Second * 2

	wg.Add(1)
	go processUntilGracefulExit(filename, totalWrites, timeout)

	for i := 0; i < totalWrites; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writeFile(filename, fmt.Sprintf("Data %d\n", i))
		}()
	}

	for i := 0; i < 5; i++ { // Example number of reads
		wg.Add(1)
		go func() {
			defer wg.Done()
			readFile(filename)
		}()
	}

	wg.Wait()
	fmt.Println("Process completed gracefully.")
}
