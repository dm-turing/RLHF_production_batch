package main

import "fmt"

func main() {
	callback := getUntrustedCallback()
	callback() // Executing untrusted code
}

func getUntrustedCallback() func() {
	// In reality, this could be a callback received from a network request or a file
	return func() {
		fmt.Println("This could be malicious code")
	}
}
