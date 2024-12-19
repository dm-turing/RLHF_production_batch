package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for _, v := range m {
		fmt.Println(v) // Output: 1 2 3
	}
}
