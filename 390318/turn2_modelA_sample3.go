package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func blocking() {
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	go blocking()

	fmt.Println("Taking goroutine profile")
	f, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile:", err)
	}
	defer f.Close()
	if err := pprof.Lookup("goroutine").WriteTo(f, 1); err != nil {
		log.Fatal("could not write goroutine profile:", err)
	}
	fmt.Println("Goroutine profile done")
}
