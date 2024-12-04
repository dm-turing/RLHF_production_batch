package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func work() {
	for i := 0; i < 1e9; i++ {
	}
}

func main() {
	fmt.Println("Starting CPU profiler")
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create cpu profile:", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start cpu profile:", err)
	}

	defer pprof.StopCPUProfile()

	// Work for 10 seconds
	time.Sleep(10 * time.Second)
	fmt.Println("CPU profiler stopped")
}
