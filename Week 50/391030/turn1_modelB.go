package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func someFunction() error {
	// original error
	_, err := os.Open("file.txt")
	if err != nil {
		return fmt.Errorf("someFunction: %w", err)
	}
	return nil
}

func divide(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func readConfig(file string) (map[string]string, error) {
	_, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %q: %w", file, err)
	}
	// Further processing...
	return nil, nil
}

func processInput(input interface{}) error {
	switch v := input.(type) {
	case int:
		fmt.Println("Input is an int:", v)
	default:
		return fmt.Errorf("unexpected input type: %T", input)
	}
	return nil
}

func main() {
	fmt.Println(fmt.Errorf("failed to open file %q: %w", "fileName", os.ErrNotExist))
	fmt.Println(someFunction())
	fmt.Println(fmt.Errorf("invalid input: expected int, got %T", "input"))
	result := divide(10, 0)
	if result == 0 {
		fmt.Println(errors.New("cannot divide by zero")) // Output: cannot divide by zero
	}
	fmt.Println(processInput("This is a string"))
}
