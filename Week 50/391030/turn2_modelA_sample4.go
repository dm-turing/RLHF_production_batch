package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func deepFunction() error {
	// Raise an error in the deepest function with a stack trace
	return errors.New("deep error occurred")
}

func middleFunction() error {
	if err := deepFunction(); err != nil {
		// Wrap the error with stack trace info
		return errors.Wrap(err, "middle function encountered an error")
	}
	return nil
}

func topFunction() error {
	if err := middleFunction(); err != nil {
		// Wrap once more for context
		return errors.Wrap(err, "top function tracking error")
	}
	return nil
}

func main() {
	err := topFunction()
	if err != nil {
		fmt.Printf("Error: %+v\n", err) // %+v prints the stack trace as well
	}
}
