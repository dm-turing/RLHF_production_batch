package main

import (
	"fmt"
	"sync"
)

func main() {
	// input := "Hello, world! Hello, Go."
	var wg sync.WaitGroup
	var wordCounts = make(map[string]int)

	// Split the input string into words
	words := []string{"Hello", "world", "Hello", "Go"}

	// Count each word using goroutines
	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			wordCounts[w] += 1
		}(word)
	}

	wg.Wait()

	// Print the results
	for word, count := range wordCounts {
		fmt.Printf("Word: %s, Count: %d\n", word, count)
	}
}
