package main

import (
	"fmt"
)

func main() {
	originalSlice := make([]string, 0, 100)
	mutableSlice := make([]string, 0, 100)

	for i := 0; i < 100; i++ {
		originalSlice = append(originalSlice, fmt.Sprintf("Element %d", i))
	}

	for _, str := range originalSlice {
		mutableSlice = append(mutableSlice, str)
	}

	// Clear references to help GC if needed
	originalSlice = nil
}
