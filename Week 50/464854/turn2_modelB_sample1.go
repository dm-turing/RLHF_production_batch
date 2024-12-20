package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Total Memory Allocated: ", m.HeapAlloc/1024/1024) // convert to MB()
}
