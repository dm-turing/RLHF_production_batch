package main

import (
	"fmt"
	"testing"
)

func BenchmarkResizeSlice(b *testing.B) {
	for _, size := range []int{10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				slice := make([]int, size)
				slice = append(slice, 1)
			}
		})
	}
}
