package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Define a high precision number
	x := big.NewFloat(0.555)

	// Set the precision
	prec := uint(2)

	// Example of rounding modes
	modes := []big.RoundingMode{
		big.ToNearestAway,
		big.ToNearestEven,
		big.ToZero,
		big.AwayFromZero,
		big.ToNegativeInf,
		big.ToPositiveInf,
	}

	for _, mode := range modes {
		z := new(big.Float).SetPrec(prec).SetMode(mode)
		z.Set(x)
		fmt.Printf("Rounding mode: %v, Result: %.2f\n", mode, z)
	}
}
