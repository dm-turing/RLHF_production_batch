package main

import (
	"math/big"
	"testing"
)

func BenchmarkExponentialOperations(b *testing.B) {
	x := big.NewFloat(2.0)
	modes := []big.RoundingMode{
		big.ToNearestAway,
		big.ToNearestEven,
	}

	for _, mode := range modes {
		b.Run(mode.String(), func(b *testing.B) {
			z := new(big.Float).SetPrec(2).SetMode(mode)
			for i := 0; i < b.N; i++ {
				z.Exp(z, x, x)
			}
		})
	}
}
