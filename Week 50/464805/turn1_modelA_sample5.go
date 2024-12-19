package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for k := range m {
			m[k]++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for k, v := range m {
			fmt.Println(k, v)
		}
	}()

	wg.Wait()
}
