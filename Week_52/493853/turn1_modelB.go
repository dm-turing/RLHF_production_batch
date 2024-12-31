package main

import "fmt"

func add(a, b int) int {
	return a + b
}

type callback func(int, int) int

func useCallback(cb callback, x, y int) int {
	return cb(x, y)
}

func main() {
	result := useCallback(add, 10, 20)
	fmt.Println(result) // Output: 30
}
