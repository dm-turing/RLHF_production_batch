package main

import (
	"debug"
	"fmt"
)

func main() {
	var m debug.MemStats
	debug.ReadMemStats(&m)
	fmt.Println("Live Slice Count: ", m.HeapSys/1024/1024) // Approximate slice count}
