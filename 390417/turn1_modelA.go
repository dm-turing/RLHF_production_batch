package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Define a floating-point number
	x, _ := new(big.Float).SetString("1.5")

	// Set rounding modes
	roundingModes := []string{
		"RoundHalfEven",
		"RoundHalfUp",
		"RoundHalfDown",
		"RoundHalfToEven",
		"RoundPlain",
	}

	for _, mode := range roundingModes {
		x.SetPrec(64) // Set precision
		x.Round(x, mode)
		fmt.Printf("Rounding Mode: %s, Result: %.6f\n", mode, x.Text('f', 6))
	}
}
