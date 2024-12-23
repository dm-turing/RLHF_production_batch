package main

import (
	"log"
	"time"
)

func logFunctionExecution(name string) {
	log.Printf("Entering function: %s\n", name)
	defer log.Printf("Exiting function: %s\n", name)
}

func timeFunctionExecution(name string) {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime)
		log.Printf("Function %s took %v\n", name, duration)
	}()
}

func sleepForSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	logFunctionExecution("main")
	timeFunctionExecution("sleepForSeconds")
	sleepForSeconds(2)
}
