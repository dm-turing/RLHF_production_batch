package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	// Example 1: Using StringBuilder for efficient concatenation
	var sb bytes.StringBuilder
	for i := 0; i < 1000000; i++ {
		sb.WriteString(fmt.Sprintf("Hello, world! Iteration: %d\n", i))
	}
	fmt.Println("StringBuilder output:")
	fmt.Println(sb.String())

	// Example 2: Using native string concatenation
	var result string
	startTime := time.Now()
	for i := 0; i < 1000000; i++ {
		result += fmt.Sprintf("Hello, world! Iteration: %d\n", i)
	}
	endTime := time.Now()
	fmt.Println("Native concatenation output:")
	fmt.Println(result)
	fmt.Printf("Native concatenation took: %s\n", endTime.Sub(startTime))

	// Example 3: Using StringBuilder with fmt.Sprintf
	var sb2 bytes.StringBuilder
	for i := 0; i < 1000000; i++ {
		sb2.WriteString(fmt.Sprintf("Hello, world! Iteration: %d\n", i))
	}
	fmt.Println("StringBuilder with fmt.Sprintf output:")
	fmt.Println(sb2.String())
}
