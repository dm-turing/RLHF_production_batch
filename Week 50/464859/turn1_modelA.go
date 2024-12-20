package main

import (
	"fmt"
	"log"
)

// ErrorHandlerDecorator takes a function f and returns a new function that logs errors
func ErrorHandlerDecorator(f func() error) func() error {
	return func() error {
		err := f()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		return err
	}
}

// Example function that may return an error
func exampleFunction() error {
	return fmt.Errorf("simulated error")
}

func main() {
	// Decorate the exampleFunction with the ErrorHandlerDecorator
	decoratedFunction := ErrorHandlerDecorator(exampleFunction)

	// Call the decorated function
	err := decoratedFunction()
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
}
