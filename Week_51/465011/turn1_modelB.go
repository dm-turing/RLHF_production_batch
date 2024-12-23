package main

import (
	"fmt"
	"sync"
)

// LazyString implements the fmt.Stringer interface, allowing lazy string
// evaluation.
type LazyString struct {
	once  sync.Once
	val   string
	build func() string
}

// NewLazyString creates a new LazyString with the specified string builder.
func NewLazyString(build func() string) LazyString {
	return LazyString{build: build}
}

// String returns the string representation of LazyString. It performs lazy
// evaluation only when called for the first time.
func (l *LazyString) String() string {
	l.once.Do(func() {
		l.val = l.build()
	})
	return l.val
}

func main() {
	// Example usage of LazyString
	type Point struct {
		X, Y int
	}

	point := Point{1, 2}
	lazyString := NewLazyString(func() string {
		return fmt.Sprintf("Point(%d, %d)", point.X, point.Y)
	})

	// Lazy evaluation happens only when the string is printed
	fmt.Println(lazyString) // Output: Point(1, 2)

	// Changes to the underlying point do not affect the lazy string
	point.X = 3
	point.Y = 4
	fmt.Println(lazyString) // Output: Point(1, 2)
}
