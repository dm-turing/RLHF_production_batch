package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for k := range m {
			m[k]++
		}
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, v)
		}
		mu.Unlock()
	}()

	wg.Wait()
}
