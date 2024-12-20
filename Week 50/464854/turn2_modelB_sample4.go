package main

import (
	"debug"
	"fmt"
)

func main() {
	var m debug.MemStats
	debug.ReadMemStats(&m)
	fmt.Println("Heap Sys: ", m.HeapSys/1024/1024)
}
