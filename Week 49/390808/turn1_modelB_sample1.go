package main

import (
	"fmt"
	"time"
)

func processItem(id int, ch chan string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Processed item %d", id)
}

func main() {
	resultCh := make(chan string, 10)

	for i := 1; i <= 10; i++ {
		go processItem(i, resultCh)
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(<-resultCh)
	}
}
