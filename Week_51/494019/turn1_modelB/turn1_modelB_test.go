package turn1_modelB

import (
	"fmt"
	"testing"
)

func BenchmarkForLoop(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	b.ResetTimer()

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	b.StopTimer()
	fmt.Println(sum)
}

func BenchmarkRangeStatement(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	b.ResetTimer()

	for _, value := range arr {
		sum += value
	}

	b.StopTimer()
	fmt.Println(sum)
}
