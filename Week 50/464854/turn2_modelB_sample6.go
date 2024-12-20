package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Escaped to Heap: ", m.HeapAlloc/1024/1024)
}
