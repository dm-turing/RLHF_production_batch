package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

// LinearCongruentialGenerator struct
type LinearCongruentialGenerator struct {
	seed, a, c, m uint32
}

// NewLCG initializes a new LCG with given parameters
func NewLCG(seed, a, c, m uint32) *LinearCongruentialGenerator {
	return &LinearCongruentialGenerator{seed: seed, a: a, c: c, m: m}
}

// Next generates the next random number
func (lcg *LinearCongruentialGenerator) Next() uint32 {
	lcg.seed = (lcg.a*lcg.seed + lcg.c) % lcg.m
	return lcg.seed
}

func main() {
	// Define flags for the maximum value and seed
	minValue := flag.Int("min", 0, "The minimum value for the random number")
	maxValue := flag.Int("max", 100, "The maximum value for the random number")
	seed := flag.Int64("seed", time.Now().UnixNano(), "Seed for the random number generator")
	flag.Parse()

	// Validate the range
	err := validateRange(*minValue, *maxValue)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// LCG parameters: a, c, and m
	a := uint32(1664525)
	c := uint32(1013904223)
	m := uint32(1 << 31) // 2^31

	// Initialize LCG
	lcg := NewLCG(uint32(*seed), a, c, m)

	// Generate a random number
	randomNumber := int(lcg.Next()%uint32(*maxValue-*minValue+1)) + *minValue
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
