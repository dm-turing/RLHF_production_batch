package main

import (
	"fmt"
	"strings"
)

// Function interface defines the signature that all callbacks must follow
type Function interface {
	Call() string
}

// StringFuncType is an anonymous function that returns a string
type StringFuncType func() string

// implements the Function interface
func (f StringFuncType) Call() string {
	return f()
}

// IntegerFuncType is an anonymous function that returns an integer
type IntegerFuncType func() int

// implements the Function interface
func (f IntegerFuncType) Call() string {
	return fmt.Sprintf("%d", f())
}

// filterCallbacks filters callbacks based on a predicate
func filterCallbacks(callbacks []Function, predicate func(string) bool) []Function {
	filteredCallbacks := make([]Function, 0)
	for _, callback := range callbacks {
		result := callback.Call()
		if predicate(result) {
			filteredCallbacks = append(filteredCallbacks, callback)
		}
	}
	return filteredCallbacks
}

// executeCallbacks calls the filtered callbacks
func executeCallbacks(callbacks []Function) {
	for _, callback := range callbacks {
		fmt.Println(callback.Call())
	}
}

func main() {
	callbacks := make([]Function, 0)

	// Register named functions
	callbacks = append(callbacks, StringFuncType(hello))
	callbacks = append(callbacks, IntegerFuncType(greet))

	// Register anonymous functions
	callbacks = append(callbacks, StringFuncType(func() string { return "Anonymous String" }))
	callbacks = append(callbacks, IntegerFuncType(func() int { return 100 }))

	// Filter callbacks that return strings starting with "H"
	filtered := filterCallbacks(callbacks, func(s string) bool {
		return strings.HasPrefix(s, "H")
	})

	// Execute the filtered callbacks
	executeCallbacks(filtered)
}

// hello is a named function returning a string
func hello() string {
	return "Hello"
}

// greet is a named function returning an integer
func greet() int {
	return 42
}
