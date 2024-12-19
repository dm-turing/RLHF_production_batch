package main

import "fmt"

func main() {
	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, val := range myMap {
		fmt.Printf("%s: %d\n", key, val)
	}
}
