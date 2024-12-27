package turn1_modelA

import (
	"testing"
)

func BenchmarkTraditionalFor(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000000; j++ {
			sum += j
		}
	}
}

func BenchmarkRangeStatement(b *testing.B) {
	var sum int
	arr := make([]int, 100000000)
	for i := 0; i < b.N; i++ {
		for _, value := range arr {
			sum += value
		}
	}
}
