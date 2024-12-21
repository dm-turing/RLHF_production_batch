package main

import (
	"fmt"
	"sync"
)

func increment(num *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		*num++
	}
}

func main() {
	var num int = 0
	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&num, &wg)
	go increment(&num, &wg)

	wg.Wait()
	fmt.Println("Final value of num:", num) // Output: Final value of num: 200000
}
