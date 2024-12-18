package main

import "fmt"

// Define a simple function type
type Adder func(int) int

// Create a function that returns an Adder
func MakeAdder(x int) Adder {
	return func(y int) int {
		return x + y
	}
}

// Use functional composition to create a new function
func MakeMultiplier(x int) Adder {
	return func(y int) int {
		return MakeAdder(x)(y) * x
	}
}

func main() {
	addFive := MakeAdder(5)
	fmt.Println(addFive(10)) // Output: 15

	multiplyByThree := MakeMultiplier(3)
	fmt.Println(multiplyByThree(5)) // Output: 30
}
