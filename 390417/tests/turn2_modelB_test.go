package main

import (
	"math/big"
	"testing"
)

// BenchmarkRoundingAndArithmetic benchmarks different rounding modes in combination with division operation.
func BenchmarkRoundingAndArithmetic(b *testing.B) {
	x := big.NewFloat(10.555)
	y := big.NewFloat(3.333)
	modes := []big.RoundingMode{
		big.ToNearestAway,
		big.ToNearestEven,
		big.ToZero,
		big.AwayFromZero,
		big.ToNegativeInf,
		big.ToPositiveInf,
	}

	for _, mode := range modes {
		b.Run(mode.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				z := new(big.Float).SetPrec(10).SetMode(mode)
				z.Quo(x, y) // Perform division
			}
		})
	}
}
