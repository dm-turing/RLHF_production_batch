package main

import (
    "fmt"
)

// Functor wraps a generic value.
type Functor[T any] struct {
    value T
}

// Map applies a function to the value inside the functor.
func (f Functor[T]) Map[U any](fn func(T) U) Functor[U] {
    return Functor[U]{value: fn(f.value)}
}

func main() {
    val := Functor[int]{value: 10}

    // Function to apply
    increment := func(x int) int { return x + 1 }

    // Use the Map function
    newVal := val.Map(increment)

    fmt.Println("New Value:", newVal.value) // Output: New Value: 11
}