package main

import (
	"fmt"
)

func square(x int) int {
	return x * x
}

func main() {
	result := square(5)
	fmt.Println(result) // Output: 25
}
