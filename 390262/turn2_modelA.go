// Package math provides basic mathematical functions.
package math

import "fmt"

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// ExampleAbs demonstrates the use of the Abs function.
func ExampleAbs() {
	fmt.Println(Abs(-5)) // Output: 5
	fmt.Println(Abs(0))  // Output: 0
	fmt.Println(Abs(7))  // Output: 7
}
