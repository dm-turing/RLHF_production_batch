package main

import "fmt"

func main() {
	secret := "highly-secret"
	f := func() {
		fmt.Println("Secret:", secret)
	}
	// Passing the closure to an untrusted function
	untrustedFunction(f)
}

func untrustedFunction(fn func()) {
	// In this scenario, an untrusted function can call the closure and access the secret
	fn()
}
