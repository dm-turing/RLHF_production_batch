package main

import (
	"fmt"
)

// Point is a struct representing a point in 2D space.
type Point struct {
	X int
	Y int
}

// String implements the Stringer interface.
func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func main() {
	p := Point{3, 4}

	// Print the point
	fmt.Println("Point:", p)

	// Using the Stringer interface in logging
	fmt.Println("Logging a point:", p)
}
