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
	"net/http"
	"time"
)

// Workload simulates a network request handler
func Workload(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "Processing request\n")
	// Simulate processing
	time.Sleep(50 * time.Millisecond)
	duration := time.Since(start)
	fmt.Fprintf(w, "Request processed in %v\n", duration)
}

func main() {
	// Set CPU affinity to core 1 for this application
	fmt.Println("Setting CPU affinity to core 1")
	C.setAffinity(1)

	http.HandleFunc("/", Workload)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
