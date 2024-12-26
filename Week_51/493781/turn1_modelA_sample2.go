package main

import (
	"fmt"
)

func main() {
	square := func(x int) int {
		return x * x
	}
	result := square(5)
	fmt.Println(result) // Output: 25
}
