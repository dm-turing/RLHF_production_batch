package main

import (
	"fmt"
)

// Point represents a point in 2D space.
type Point struct {
	X int
	Y int
}

// Format implements the fmt.Formatter interface to provide custom formatting.
func (p Point) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		// Default format for verb 'v'
		fmt.Fprintf(f, "%+v\n", p)
	case 's':
		// Custom format for verb 's'
		fmt.Fprintf(f, "(%d, %d)\n", p.X, p.Y)
	default:
		// Handle unknown verbs
		fmt.Fprintf(f, "unknown verb %c\n", verb)
	}
}

func main() {
	p := Point{3, 4}

	// Default verb 'v'
	fmt.Println(p)

	// Custom string verb 's'
	fmt.Printf("%s\n", p)

	// You can also use with %q for quoted string
	fmt.Printf("%q\n", p)

	// Custom verb 'x' (not implemented here, will use default format)
	fmt.Printf("%x\n", p)
}
