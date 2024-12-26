package main

import "fmt"

func main() {
	d := Data{Secret: "highly-secret"}
	// Passing a method of a struct instead of a closure
	untrustedFunction(d.Print)
}

func untrustedFunction(fn func()) {
	fn()
}

type Data struct {
	Secret string
}

func (d Data) Print() {
	fmt.Println("Secret:", d.Secret)
}
