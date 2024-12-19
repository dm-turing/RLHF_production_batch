// Future Trend: Parallel iteration with range
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(num int) {
			fmt.Println(num)
			wg.Done()
		}(num)
	}

	wg.Wait()
}
