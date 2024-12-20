package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Live Slice Count: ", m.HeapSys/1024/1024) // Approximate slice count
}
