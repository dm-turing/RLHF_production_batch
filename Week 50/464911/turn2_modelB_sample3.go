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

// Define an applicative for slices
type SliceApplicative []func(int) int

func (a SliceApplicative) Apply(f SliceFunctor) SliceFunctor {
	result := make(SliceFunctor, len(f))
	for i, x := range f {
		for _, g := range a {
			result[i] = g(x)
		}
	}
	return result
}

func main() {
	functions := SliceApplicative{func(x int) int { return x * 2 }, func(x int) int { return x + 1 }}
	numbers := SliceFunctor{1, 2, 3, 4, 5}
	results := functions.Apply(numbers)
	fmt.Println(results) // Output: [2 3 4 5 6 3 4 5 6 7]
}
