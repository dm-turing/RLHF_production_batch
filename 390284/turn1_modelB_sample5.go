package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello(name string) {
	fmt.Println("Hello from", name, ", GOMAXPROCS:", runtime.GOMAXPROCS(0))
	time.Sleep(1 * time.Second)
}

func main() {
	runtime.GOMAXPROCS(2) // Set GOMAXPROCS to 2

	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")
	go sayHello("Goroutine 3")

	fmt.Println("Main goroutine is done.")
}
