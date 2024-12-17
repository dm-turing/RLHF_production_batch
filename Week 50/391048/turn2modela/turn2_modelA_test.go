package main

import (
	"fmt"
	"testing"
)

const size = 100000

// Benchmark for Slices
func BenchmarkSliceAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slice := make([]int, 0, size)
		for i := 0; i < size; i++ {
			slice = append(slice, i)
		}
	}
}

// Benchmark for Maps
func BenchmarkMapInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m := make(map[int]int, size)
		for i := 0; i < size; i++ {
			m[i] = i
		}
	}
}

// Benchmark for Slices iteration
func BenchmarkSliceIteration(b *testing.B) {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}

	b.ResetTimer() // Reset timer for actual benchmarking
	for n := 0; n < b.N; n++ {
		for _, v := range slice {
			_ = v // Just to prevent compiler optimization
		}
	}
}

// Benchmark for Maps iteration
func BenchmarkMapIteration(b *testing.B) {
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i
	}

	b.ResetTimer() // Reset timer for actual benchmarking
	for n := 0; n < b.N; n++ {
		for _, v := range m {
			_ = v // Just to prevent compiler optimization
		}
	}
}

func main() {
	// Run the benchmarks using `go test` command
	fmt.Println("Run 'go test -bench=.' to see benchmark results.")
}
