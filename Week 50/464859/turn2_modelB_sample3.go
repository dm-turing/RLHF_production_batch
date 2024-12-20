package main

import (
	"fmt"
	"log"
)

type Calculator interface {
	Add(int, int) int
}

type LoggingCalculator struct {
	Calculator
}

func (lc *LoggingCalculator) Add(a, b int) int {
	result := lc.Calculator.Add(a, b)
	log.Printf("Calculated: %d + %d = %d\n", a, b, result)
	return result
}

func NewLoggingCalculator(calc Calculator) *LoggingCalculator {
	return &LoggingCalculator{calc}
}

type RealCalculator struct{}

func (rc *RealCalculator) Add(a, b int) int {
	return a + b
}

func main() {
	calc := &RealCalculator{}
	loggingCalc := NewLoggingCalculator(calc)
	result := loggingCalc.Add(10, 20)
	fmt.Println("Result:", result)
}
