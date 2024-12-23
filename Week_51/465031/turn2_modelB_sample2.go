package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("large_text_file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wordCount := make(map[string]*sync.Mutex)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			mutex, _ := wordCount[word]
			if mutex == nil {
				mutex = &sync.Mutex{}
				wordCount[word] = mutex
			}
			mutex.Lock()
			wordCount[word] = wordCount[word] + 1
			mutex.Unlock()
		}
	}

	// Print the word count
	fmt.Println("Word Count:")
	for word, count := range wordCount {
		count.Lock()
		fmt.Printf("%s: %d\n", word, count)
		count.Unlock()
	}
}
