package main

import (
	"fmt"
	"sync"
)

func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func mergeMapsParallel(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	var wg sync.WaitGroup

	wg.Add(len(maps))

	for _, m := range maps {
		go func(m map[string]int) {
			for k, v := range m {
				result[k] = v
			}
			wg.Done()
		}(m)
	}

	wg.Wait()

	return result
}

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	map2 := map[string]int{
		"d": 4,
		"e": 5,
		"f": 6,
	}

	map3 := map[string]int{
		"g": 7,
		"h": 8,
		"i": 9,
	}

	result := mergeMaps(map1, map2, map3)
	fmt.Println("Result:", result)

	result = mergeMapsParallel(map1, map2, map3)
	fmt.Println("Result (parallel):", result)
}
