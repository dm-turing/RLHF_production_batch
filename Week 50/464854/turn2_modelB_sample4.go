package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Heap Sys: ", m.HeapSys/1024/1024)
}
