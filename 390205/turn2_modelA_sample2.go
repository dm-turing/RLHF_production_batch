package main

import (
	"fmt"
)

type MyExpensiveType struct {
	// Replace with your data structure
}

var lazyValue *MyExpensiveType

func initLazy() *MyExpensiveType {
	if lazyValue == nil {
		lazyValue = &MyExpensiveType{}
	}
	return lazyValue
}

func main() {
	for i := 0; i < 10; i++ {
		// Use initLazy() concurrently
		fmt.Println("Getting lazyValue from goroutine", i)
		fmt.Println("Value:", initLazy())
	}
}
