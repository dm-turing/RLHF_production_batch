package main

import (
	"fmt"
	"log"
)

// ErrorHandlerDecoratorMultiple takes a function f and returns a new function that logs errors
func ErrorHandlerDecoratorMultiple(f func(string) (int, error)) func(string) (int, error) {
	return func(input string) (int, error) {
		result, err := f(input)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		return result, err
	}
}

// Example function that may return an error
func exampleFunctionMultiple(input string) (int, error) {
	if input == "" {
		return 0, fmt.Errorf("input is empty")
	}
	return len(input), nil
}

func main() {
	// Decorate the exampleFunctionMultiple with the ErrorHandlerDecoratorMultiple
	decoratedFunction := ErrorHandlerDecoratorMultiple(exampleFunctionMultiple)

	// Call the decorated function
	result, err := decoratedFunction("Hello")
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
	fmt.Printf("Result: %d\n", result)

	// Call the decorated function with an empty input to simulate an error
	_, err = decoratedFunction("")
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
}
