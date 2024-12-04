package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	var unfreed []byte

	for i := 0; i < 100; i++ {
		unfreed = append(unfreed, []byte(fmt.Sprintf("data%d", i))...)
		fmt.Println("Allocated", len(unfreed), "bytes")
	}

	// Snapshot memory usage
	fmt.Println("Taking memory profile")
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile:", err)
	}
	defer f.Close()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile:", err)
	}
	fmt.Println("Memory profile done")
}
