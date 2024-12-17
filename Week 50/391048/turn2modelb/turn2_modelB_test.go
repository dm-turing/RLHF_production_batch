package main

import (
	"benchmark"
	"fmt"
	"testing"
)

// Benchmark for slices
func BenchmarkSliceAppend(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
}

func BenchmarkSliceIteration(b *testing.B) {
	s := make([]int, b.N)
	for i := range s {
		s[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s[i]
	}
}

// Benchmark for maps
func BenchmarkMapInsert(b *testing.B) {
	var m map[int]int
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkMapLookup(b *testing.B) {
	m := make(map[int]int, b.N)
	for i := range m {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func BenchmarkMapDelete(b *testing.B) {
	m := make(map[int]int, b.N)
	for i := range m {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(m, i)
	}
}

func main() {
	fmt.Println("Running benchmarks...")
	benchmark.Run()
}
