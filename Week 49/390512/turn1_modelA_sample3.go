package main

import "fmt"

func divide(x, y int) int {
	if y == 0 {
		panic("division by zero") // Critical logic error
	}
	return x / y
}

func main() {
	result := divide(10, 0) // This will panic
	fmt.Println(result)     // This line won't execute
}
