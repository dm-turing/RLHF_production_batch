package main

import (
	"math/big"
	"testing"
)

func BenchmarkBasicOperations(b *testing.B) {
	x := big.NewFloat(0.555)
	y := big.NewFloat(0.222)
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
			z := new(big.Float).SetPrec(2).SetMode(mode)
			for i := 0; i < b.N; i++ {
				z.Add(x, y)
			}
		})
	}
}
