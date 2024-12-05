package main

import (
	"math/big"
	"testing"
)

func BenchmarkRoundHalfEven(b *testing.B) {
	x, _ := new(big.Float).SetString("1.5")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x.SetPrec(64)
		x.Round(x, "RoundHalfEven")
	}
}

func BenchmarkRoundHalfUp(b *testing.B) {
	x := new(big.Float).SetString("1.5")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x.SetPrec(64)
		x.Round(x, "RoundHalfUp")
	}
}

// Add more benchmark functions for other rounding modes
