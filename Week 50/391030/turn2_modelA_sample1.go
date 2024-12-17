package main

import (
	"errors"
	"fmt"
)

func deepestFunction() error {
	// Simulating an error
	return errors.New("deep error")
}

func middleFunction() error {
	if err := deepestFunction(); err != nil {
		return fmt.Errorf("middleFunction failed: %w", err)
	}
	return nil
}

func topFunction() error {
	if err := middleFunction(); err != nil {
		return fmt.Errorf("topFunction failed: %w", err)
	}
	return nil
}

func main() {
	err := topFunction()
	if err != nil {
		fmt.Println(err)
		// You can use errors.Unwrap to get to the original error
		if unwrapped := errors.Unwrap(err); unwrapped != nil {
			fmt.Println("Unwrapped error:", unwrapped)
		}
	}
}
