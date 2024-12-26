package main

import (
	"fmt"
)

func calculateSum(a int, b int) int {
	return a + b
}

func funcExample() {
	counter := 0

	namedFunction := func() {
		counter += 1 // This will not affect the outer counter
	}

	anonymousFunction := func() {
		counter += 1 // This will modify the outer counter
	}

	namedFunction()
	anonymousFunction()
	fmt.Println(counter) // Output: 1
}

func main() {
	fmt.Println(calculateSum(10, 20))
	funcExample()
}
