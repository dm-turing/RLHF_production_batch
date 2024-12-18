package main

import "fmt"

func SomeOtherFunction() error {
	return fmt.Errorf("%s", "Some error happened")
}

func SomeFunction() error {
	// Perform some operation
	if err := SomeOtherFunction(); err != nil {
		return fmt.Errorf("SomeFunction failed: %w", err)
	}
	return nil
}

func main() {
	if err := SomeFunction(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success!")
	}
}
