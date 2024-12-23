package main

import (
	"fmt"
)

// LazyString is a closure that returns a string when called.
type LazyString func() string

// String implements the Stringer interface.
func (ls LazyString) String() string {
	return ls()
}

func main() {
	// Lazy string example
	var l LazyString = func() string {
		fmt.Println("Computing string...")
		return "Lazy String Representation"
	}

	// String is not computed until it's printed
	fmt.Println("Non-lazy string:", "Hello, World!")
	fmt.Println("Lazy string:", l)
}
