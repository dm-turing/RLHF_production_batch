package main

import (
	"testing"
)

func BenchmarkCopyByValue(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		copyByValue(slice)
	}
}

func copyByValue(slice []int) []int {
	return append([]int{}, slice...)
}

func BenchmarkCopyByReference(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		copyByReference(&slice)
	}
}

func copyByReference(slice *[]int) {
	*slice = append([]int{}, *slice...)
}

func BenchmarkResizeByValue(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		resizeByValue(slice)
	}
}

func resizeByValue(slice []int) {
	slice = append(slice, 100)
}

func BenchmarkResizeByReference(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		resizeByReference(&slice)
	}
}

func resizeByReference(slice *[]int) {
	*slice = append(*slice, 100)
}

func BenchmarkAppendByValue(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		appendByValue(slice)
	}
}

func appendByValue(slice []int) {
	slice = append(slice, 100)
}

func BenchmarkAppendByReference(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		appendByReference(&slice)
	}
}

func appendByReference(slice *[]int) {
	*slice = append(*slice, 100)
}
