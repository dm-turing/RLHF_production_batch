package main

import "fmt"

// Define a functor for slices
type SliceFunctor []int

func (f SliceFunctor) Map(fn func(int) int) SliceFunctor {
	result := make(SliceFunctor, len(f))
	for i, x := range f {
		result[i] = fn(x)
	}
	return result
}

func main() {
	numbers := SliceFunctor{1, 2, 3, 4, 5}
	doubled := numbers.Map(func(x int) int { return x * 2 })
	fmt.Println(doubled) // Output: [2 4 6 8 10]
}
