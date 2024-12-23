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

	var wordCount sync.Map // Using sync.Map
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			v, _ := wordCount.LoadOrStore(word, 0)
			wordCount.Store(word, v.(int)+1)
		}
	}

	// Print the word count
	fmt.Println("Word Count:")
	wordCount.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", key, value.(int))
		return true
	})
}
