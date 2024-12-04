package main

import (
	"fmt"
)

// ArrayFunctor represents a functor for arrays.
type ArrayFunctor struct {
	slice []int
}

func (f ArrayFunctor) Map(fn func(int) int) ArrayFunctor {
	return ArrayFunctor{slice: f.applyFunc(fn)}
}

func (f ArrayFunctor) applyFunc(fn func(int) int) []int {
	out := make([]int, len(f.slice))
	for i, x := range f.slice {
		out[i] = fn(x)
	}
	return out
}

func main() {
	arr := ArrayFunctor{slice: []int{1, 2, 3, 4, 5}}

	doubledArr := arr.Map(func(x int) int { return x * 2 })
	fmt.Println("Doubled array:", doubledArr.slice)

	squareArr := arr.Map(func(x int) int { return x * x })
	fmt.Println("Squared array:", squareArr.slice)
}
