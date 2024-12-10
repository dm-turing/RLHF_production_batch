package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Metrics for I/O operation type distribution
	ioOperationCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "io_operation_count",
			Help: "Total number of I/O operations by type.",
		},
		[]string{"operation_type"},
	)

	// Metrics for resource utilization
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage_percentage",
		Help: "Current CPU usage percentage.",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage_percentage",
		Help: "Current memory usage percentage.",
	})
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(ioOperationCount)
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
}

func recordMetrics() {
	go func() {
		for {
			// Simulating I/O operations
			ioOperationCount.WithLabelValues("read").Add(float64(rand.Intn(100)))
			ioOperationCount.WithLabelValues("write").Add(float64(rand.Intn(100)))
			ioOperationCount.WithLabelValues("delete").Add(float64(rand.Intn(100)))

			// Simulating resource usage metrics
			cpuUsage.Set(float64(rand.Intn(100)))    // Random CPU usage
			memoryUsage.Set(float64(rand.Intn(100))) // Random memory usage

			time.Sleep(5 * time.Second) // Scrape every 5 seconds
		}
	}()
}

func main() {
	// Start recording metrics
	recordMetrics()

	// Serve the /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
