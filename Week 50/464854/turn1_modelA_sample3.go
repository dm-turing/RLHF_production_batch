package main

func main() {
	data := make(map[int][]int)
	for i := 0; i < 1000000; i++ {
		largeSlice := make([]int, i)
		data[i] = largeSlice
	}
	// If keys are lost and slices are not removed from the map, memory is retained
}
