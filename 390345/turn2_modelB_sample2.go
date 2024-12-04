package main

import (
    "fmt"
)

// Maybe represents a simple Maybe monad in Go.
type Maybe[T any] struct {
    value *T
}

// Just creates a Maybe with a value.
func Just[T any](val T) Maybe[T] {
    return Maybe[T]{value: &val}
}

// Nothing creates a Maybe with no value.
func Nothing[T any]() Maybe[T] {
    return Maybe[T]{value: nil}
}

// FlatMap applies a function and flattens the result.
func (m Maybe[T]) FlatMap[U any](fn func(T) Maybe[U]) Maybe[U] {
    if m.value == nil {
        return Nothing[U]()
    }
    return fn(*m.value)
}

// Map applies a function to the value inside the maybe.
func (m Maybe[T]) Map[U any](fn func(T) U) Maybe[U] {
    if m.value == nil {
        return Nothing[U]()
    }
    return Just(fn(*m.value))
}

func main() {
    // Example usage of Maybe monad
    maybeValue := Just(5)

    increment := func(x int) Maybe[int] {
        return Just(x + 1)
    }

    result := maybeValue.FlatMap(increment).FlatMap(increment)
    if result.value != nil {
        fmt.Println("Result:", *result.value) // Output: Result: 7
    } else {
        fmt.Println("Result: Nothing")
    }
}