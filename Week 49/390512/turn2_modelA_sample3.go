package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	side float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return 4 * s.side
}

func getArea(shape Shape) float64 {
	circle, ok := shape.(Circle)
	if !ok {
		panic("expected Circle type") // Critical: we must have a Circle to calculate area
	}
	return circle.Area()
}

func main() {
	var shape Shape = Circle{Radius: 5}

	fmt.Println("Area:", getArea(shape)) // Works fine

	var anotherShape Shape = Square{side: 5}

	// This will panic
	fmt.Println("Area:", getArea(anotherShape))
}
