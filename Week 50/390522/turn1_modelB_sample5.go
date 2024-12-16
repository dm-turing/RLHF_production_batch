package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func callMethod(shape Shape) {
	v := reflect.ValueOf(shape)
	method := v.MethodByName("Area")
	if method.IsValid() {
		result := method.Call(nil)
		fmt.Println("Area:", result[0].Interface())
	} else {
		fmt.Println("Method not found")
	}
}

func main() {
	square := Square{Side: 5.0}
	callMethod(square) // Output: Area: 25
}
