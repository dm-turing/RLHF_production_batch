package main

import (
	"fmt"
	"sync"
)

// Option emulates the option monad.
type Option int

func (o Option) IsSome() bool {
	return o != 0
}

func (o Option) IsNone() bool {
	return !o.IsSome()
}

func (o Option) Map(fn func(int) int) Option {
	if !o.IsSome() {
		return 0
	}
	return Option(fn(int(o)))
}

func (o Option) Chain(fn func(int) Option) Option {
	return fn(int(o))
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	resultChan := make(chan Option)

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result := n * 2
			if result > 0 {
				resultChan <- Option(result)
			} else {
				resultChan <- 0 // None
			}
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := []Option{}
	for result := range resultChan {
		results = append(results, result)
	}

	fmt.Println("Results:", results)

	someSum := 0
	for _, r := range results {
		if r.IsSome() {
			someSum += int(r)
		}
	}

	fmt.Println("Sum of Some values:", someSum)
}
