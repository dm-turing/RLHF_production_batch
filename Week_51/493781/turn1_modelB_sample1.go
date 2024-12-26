package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(10, 20)
	fmt.Println(result) // Output: 30
}
