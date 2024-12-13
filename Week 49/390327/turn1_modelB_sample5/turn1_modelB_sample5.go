package main

import "fmt"

var logFunc = fmt.Println

func doSomething() {
	logFunc("Doing something")
}

func main() {
	doSomething()
}
