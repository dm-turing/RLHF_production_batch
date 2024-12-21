package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

const (
	numWorkers = 3
	chunkSize  = 10
)

func uppercase(input <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range input {
		output <- strings.ToUpper(line)
	}
}
func reverse(input <-chan string, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range input {
		var reversed string
		for _, r := range line {
			reversed = string(r) + reversed
		}
		output <- reversed
	}
}
func println(input <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range input {
		fmt.Println(line)
	}
}
func main() {
	// Create multiple pipes to connect the stages
	pr, pw := io.Pipe()
	defer pw.Close()

	pr2, pw2 := io.Pipe()
	defer pw2.Close()

	pr3, _ := io.Pipe()
	defer pr3.Close()

	var wg sync.WaitGroup

	// Stage 1: Uppercase words
	wg.Add(1)
	go uppercase(chanFromPipe(pr), chanToPipe(pw), &wg)

	// Stage 2: Reverse words
	wg.Add(1)
	go reverse(chanFromPipe(pr2), chanToPipe(pw2), &wg)

	// Stage 3: Print the reversed and uppercased words
	wg.Add(1)
	go println(chanFromPipe(pr3), &wg)

	// Write data to the first pipe
	go func() {
		defer pw.Close()
		data := []string{"hello", "world", "golang", "is", "awesome"}
		for _, line := range data {
			_, err := pw.Write([]byte(line + "\n"))
			if err != nil {
				fmt.Println("Error writing to pipe:", err)
				return
			}
		}
	}()

	// Connect the pipes
	go io.Copy(pw2, pr)
	go io.Copy(pw, pr2)
	go io.Copy(pr3, pw2)

	wg.Wait()
}

// Helper functions to convert between pipes and channels
func chanFromPipe(r *io.PipeReader) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("Error reading from pipe:", err)
				return
			}
			ch <- strings.TrimSuffix(line, "\n")
		}
	}()
	return ch
}

func chanToPipe(w *io.PipeWriter) chan<- string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for line := range ch {
			n, err := w.Write([]byte(line + "\n"))
			if err != nil {
				fmt.Println("Error writing to pipe:", err)
				return
			}
			if n != len(line)+1 {
				fmt.Println("Short write to pipe:", n, "bytes written")
				return
			}
		}
	}()
	return ch
}
