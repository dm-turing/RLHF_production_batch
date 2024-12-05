package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define flags for the minimum and maximum values
	minValue := flag.Int("min", 0, "The minimum value for the random number")
	maxValue := flag.Int("max", 100, "The maximum value for the random number")
	flag.Parse()

	// Validate the range
	err := validateRange(*minValue, *maxValue)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate the random number
	randomNumber := rand.Intn(*maxValue-*minValue+1) + *minValue
	fmt.Printf("Random number between %d and %d: %d\n", *minValue, *maxValue, randomNumber)
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
