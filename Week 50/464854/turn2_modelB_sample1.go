package main

import (
	"debug"
	"fmt"
)

func main() {
	var m debug.MemStats
	debug.ReadMemStats(&m)
	fmt.Println("Total Memory Allocated: ", m.HeapAlloc/1024/1024) // convert to MB()
}
