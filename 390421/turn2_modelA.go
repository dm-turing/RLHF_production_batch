package main

/*
#include <sched.h>
#include <pthread.h>

void setAffinity(int cpu) {
    cpu_set_t cpus;
    CPU_ZERO(&cpus);
    CPU_SET(cpu, &cpus);
    pthread_setaffinity_np(pthread_self(), sizeof(cpu_set_t), &cpus);
}
*/
import "C"
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Setting CPU affinity to core 0 for the main thread")
	C.setAffinity(0)

	go func() {
		fmt.Println("Running goroutine on core 0")
		for i := 0; i < 10; i++ {
			fmt.Printf("Iteration %d\n", i)
			time.Sleep(1 * time.Second)
		}
	}()

	select {} // run indefinitely
}
