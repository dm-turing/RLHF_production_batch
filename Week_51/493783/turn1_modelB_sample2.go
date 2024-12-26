package main

import "fmt"

func apply(f func(int) int, x int) int {
	return f(x)
}

func main() {
	// Anonymous function as a callback
	result := apply(func(n int) int {
		return n * 2
	}, 10)
	fmt.Println(result) // Output: 20
}
