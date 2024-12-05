package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Set the maximum number of CPUs that can execute Go code simultaneously
	runtime.GOMAXPROCS(4)

	fmt.Println("GOMAXPROCS set to", runtime.GOMAXPROCS(0))
}
