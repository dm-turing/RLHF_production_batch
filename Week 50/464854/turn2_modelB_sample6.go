package main

import (
	"debug"
	"fmt"
)

func main() {
	var m debug.MemStats
	debug.ReadMemStats(&m)
	fmt.Println("Escaped to Heap: ", m.HeapAlloc/1024/1024)
}
