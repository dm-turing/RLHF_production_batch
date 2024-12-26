package main

import "fmt"

func main() {
	secretData := "This is a secret"
	doSomething(func() {
		fmt.Println(secretData) // Accessing sensitive data from the parent scope
	})
}

func doSomething(callback func()) {
	// Some operations here...
	callback()
}
