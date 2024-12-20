package main

import (
	"fmt"
	"log"
)

func decorateMultipleReturnValues(f func() (int, error)) func() (int, error) {
	return func() (int, error) {
		result, err := f()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		return result, err
	}
}

func someFunction() (int, error) {
	// Some business logic
	// ...
	return 42, fmt.Errorf("some error")
}

func anotherDecorator(f func() (int, error)) func() (int, error) {
	return func() (int, error) {
		result, err := f()
		if err != nil {
			return -1, fmt.Errorf("prepended error: %w", err)
		}
		return result, nil
	}
}

func main() {
	decoratedSomeFunction := anotherDecorator(decorateMultipleReturnValues(someFunction))
	result, err := decoratedSomeFunction()
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
	fmt.Println("Result:", result)
}
