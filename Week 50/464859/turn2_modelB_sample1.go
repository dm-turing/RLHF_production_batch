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

func main() {
	decoratedSomeFunction := decorateMultipleReturnValues(someFunction)
	result, err := decoratedSomeFunction()
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
	fmt.Println("Result:", result)
}
