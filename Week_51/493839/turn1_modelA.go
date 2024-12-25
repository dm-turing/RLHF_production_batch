package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Example 1: Using StringBuilder for efficient concatenation
	var sb strings.Builder
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		sb.WriteString(fmt.Sprintf("Hello, world! Iteration: %d\n", i))
	}
	endTime := time.Now()
	fmt.Printf("strings.Builder concatenation took: %s\n", endTime.Sub(startTime))

	// Example 2: Using native string concatenation
	var result string
	startTime = time.Now()
	for i := 0; i < 1000; i++ {
		result += fmt.Sprintf("Hello, world! Iteration: %d\n", i)
	}
	endTime = time.Now()
	fmt.Printf("Native concatenation took: %s\n", endTime.Sub(startTime))

}
