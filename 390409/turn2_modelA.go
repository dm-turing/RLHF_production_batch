package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// lcg initializes and generates random numbers using the Linear Congruential Generator algorithm
func lcg(min, max int, seed int) ([]int, error) {
	// Constants for the LCG algorithm
	a := 1664525
	c := 1013904223
	m := int(1e9 + 9)

	if seed < 0 || seed >= m {
		return nil, fmt.Errorf("seed must be in the range [0, %d]", m-1)
	}

	if min >= max {
		return nil, fmt.Errorf("minimum value must be less than maximum value")
	}

	randomNumbers := make([]int, 0)
	current := seed

	for len(randomNumbers) < 10 { // Generate at most 10 random numbers
		current = (a*current + c) % m
		randomNumber := (current % (max - min + 1)) + min
		if !isInSet(randomNumber, randomNumbers) { // Ensure uniqueness
			randomNumbers = append(randomNumbers, randomNumber)
		}
	}

	return randomNumbers, nil
}

// isInSet checks if a number is already in the set
func isInSet(number int, set []int) bool {
	for _, n := range set {
		if n == number {
			return true
		}
	}
	return false
}

func main() {
	// Define flags for the CLI
	minValue := flag.Int("min", 0, "The minimum value for the random numbers")
	maxValue := flag.Int("max", 100, "The maximum value for the random numbers")
	seed := flag.Int("seed", 0, "The seed for the random number generator")
	flag.Parse()

	// Validate the range
	if err := validateRange(*minValue, *maxValue); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate random numbers
	randomNumbers, err := lcg(*minValue, *maxValue, *seed)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output random numbers
	fmt.Printf("Generated random numbers (seed: %d): %v\n", *seed, randomNumbers)
}

// validateRange checks if the min and max values are valid
func validateRange(min, max int) error {
	if min > max {
		return errors.New("minimum value cannot be greater than maximum value")
	}
	if min < 0 || max < 0 {
		return errors.New("values cannot be negative")
	}
	return nil
}
