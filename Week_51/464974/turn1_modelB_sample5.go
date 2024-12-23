package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution time:", time.Since(start))
	}()
	// Perform some operations here
	time.Sleep(5 * time.Second)
}
