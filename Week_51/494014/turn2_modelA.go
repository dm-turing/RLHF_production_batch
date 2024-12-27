package main

import (
	"fmt"
)

// mergeMaps merges multiple maps with preallocated capacity
func mergeMaps(maps ...map[string]int) map[string]int {
	totalCapacity := 0
	for _, m := range maps {
		totalCapacity += len(m)
	}

	// Create a new map with preallocated capacity
	mergedMap := make(map[string]int, totalCapacity)

	// Merge each map into the new map
	for _, m := range maps {
		for key, value := range m {
			mergedMap[key] = value
		}
	}

	return mergedMap
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 20, "d": 4}
	map3 := map[string]int{"e": 5, "f": 6}

	// Merge the maps with preallocated capacity
	merged := mergeMaps(map1, map2, map3)

	// Print the merged map
	fmt.Println("Merged Map:", merged)
}
