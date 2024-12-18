package main

import "fmt"

// Define a monoid for strings
type StringMonoid string

func (m StringMonoid) Append(n StringMonoid) StringMonoid {
	return m + n
}

func (m StringMonoid) Identity() StringMonoid {
	return ""
}

func main() {
	strings := []StringMonoid{"hello", " ", "world"}
	result := strings[0]
	for _, s := range strings[1:] {
		result = result.Append(s)
	}
	fmt.Println(result) // Output: hello world
}
