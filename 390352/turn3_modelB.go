package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// I/O operation counters
	ioReadOps = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "io_read_operations_total",
		Help: "Total number of read operations",
	})
	ioWriteOps = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "io_write_operations_total",
		Help: "Total number of write operations",
	})

	// Resource utilization gauges
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage_percent",
		Help: "CPU usage percentage",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage_percent",
		Help: "Memory usage percentage",
	})
)

func main() {
	// Register metrics with Prometheus
	prometheus.MustRegister(ioReadOps)
	prometheus.MustRegister(ioWriteOps)
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)

	// Start HTTP server to expose metrics
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting metrics server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func simulateIOOperations() {
	rand.Seed(time.Now().UnixNano())

	for {
		opType := rand.Intn(2) // 0 for read, 1 for write
		if opType == 0 {
			ioReadOps.Inc()
		} else {
			ioWriteOps.Inc()
		}

		time.Sleep(1 * time.Second)
	}
}

func simulateResourceUtilization() {
	for {
		// Simulate random CPU and memory usage
		cpuUsage.Set(float64(rand.Intn(100)))
		memoryUsage.Set(float64(rand.Intn(100)))

		time.Sleep(5 * time.Second)
	}
}
