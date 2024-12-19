package main

import (
	"fmt"
)

func main() {
	var originalSlice = make([]string, 0, 100)
	var mutableSlice = make([]string, 0, 100)

	for i := 0; i < 100; i++ {
		originalSlice = append(originalSlice, fmt.Sprintf("Element %d", i))
	}

	for _, str := range originalSlice {
		mutableSlice = append(mutableSlice, str)
	}

	// Assuming `originalSlice` is not referenced elsewhere, it will be GCed once done.
}
