// Race condition in Go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func increment(num *int) {
	for i := 0; i < 100000; i++ {
		*num++
	}
}

func main() {
	var num int = 0
	runtime.GOMAXPROCS(runtime.NumCPU())

	go increment(&num)
	go increment(&num)

	time.Sleep(time.Second)
	fmt.Println("Final value of num:", num) // Output: Final value of num: 171345
}
