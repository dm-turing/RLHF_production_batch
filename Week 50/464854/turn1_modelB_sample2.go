package main

func main() {
	expectedSize := 10
	// Example of bounded slice append
	slice := make([]int, 0, expectedSize) // preallocate with capacity
	for i := 0; i < expectedSize; i++ {
		slice = append(slice, i)
	}
}
