package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Used to create new StringBuilder with size 10  and capacity 20
	sb := strings.Builder{
		Size: 10,
	}

	name := "akash"
	age := 19

	// Using fmt.Sprintf to format a string
	result := fmt.Sprintf("Hello, my name is %s and I am %d years old", name, age)
	fmt.Println("Output using fmt.Sprintf: ", result)

	//Using StringBuilder to join multiple strings
	sb.WriteString("Hello, ")
	sb.WriteString("my name is ")
	sb.WriteString(name)
	sb.WriteString(" and I am ")
	fmt.Println("Output using StringBuilder: ", sb.String())

	iterativeAppend(1000000)
	optimalAppend(1000000)
}

// Iterative approach to concatenate strings
func iterativeAppend(n int) {
	start := time.Now()
	var result string
	for i := 0; i < n; i++ {
		result += fmt.Sprintf("String %d ", i)
	}
	fmt.Println("Time taken for Iterative Append:", time.Since(start))
}

// Optimal approach using StringBuilder for concatenation
func optimalAppend(n int) {
	start := time.Now()
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(fmt.Sprintf("String %d ", i))
	}
	fmt.Println("Time taken for Optimal Append:", time.Since(start))
}
