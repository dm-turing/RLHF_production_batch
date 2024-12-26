package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}
func main() {
	result := add(10, 20)
	fmt.Println("Sum of 10 and 20 is", result)
}
