package main

import (
	"fmt"
	"sync"
)

func multiplyMatrices(matrixA, matrixB [][]int, result [][]int, rowStart, rowEnd int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := rowStart; i < rowEnd; i++ {
		for j := 0; j < len(matrixB[0]); j++ {
			for k := 0; k < len(matrixB); k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
}

func main() {
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrixB := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	result := make([][]int, len(matrixA))
	for i := range len(matrixA) {
		result[i] = make([]int, len(matrixA))
	}

	numGoroutines := 3
	wg := &sync.WaitGroup{}

	for goroutineIndex := 0; goroutineIndex < numGoroutines; goroutineIndex++ {
		rowStart := (goroutineIndex * len(matrixA)) / numGoroutines
		rowEnd := ((goroutineIndex + 1) * len(matrixA)) / numGoroutines
		wg.Add(1)
		go multiplyMatrices(matrixA, matrixB, result, rowStart, rowEnd, wg)
	}

	wg.Wait()

	fmt.Println("Result:")
	for _, row := range result {
		fmt.Println(row)
	}
}
