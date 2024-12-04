package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
	responseTimeHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:         "http_response_duration_seconds",
			Help:         "Response time of HTTP requests",
			BucketWidths: []float64{0.1, 0.2, 0.5, 1, 2, 5, 10, 20, 50, 100, 200},
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(responseTimeHistogram)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path

	requestCounter.WithLabelValues(method, path).Inc()

	start := time.Now()

	// Simulate some work
	time.Sleep(time.Millisecond * 50)

	responseTimeHistogram.WithLabelValues(method, path).Observe(time.Since(start).Seconds())

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/metrics", promhttp.Handler())
	http.HandleFunc("/", handleRequest)

	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	// Periodically collect and publish runtime statistics
	go func() {
		for {
			time.Sleep(time.Minute)
			var gcStats runtime.GCStats
			runtime.ReadGCStats(&gcStats)
			prometheus.SetGauge(prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Name: "gc_pause_time_nanoseconds",
					Help: "Total time spent in GC.",
				},
				[]string{"generation"},
			).WithLabelValues("all"), float64(gcStats.PauseTotalNs))
		}
	}()

	log.Fatal(http.ListenAndServe(":9090", nil))
}
