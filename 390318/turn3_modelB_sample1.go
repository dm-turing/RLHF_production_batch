package main

import "github.com/prometheus/client_golang/prometheus"

var requestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name: "request_duration_seconds",
        Help: "Histogram of request durations",
        Buckets: prometheus.DefBuckets,
    },
    []string{"path"},
)

func init() {
    prometheus.MustRegister(requestDuration)
}

// Usage in a handler
start := time.Now()
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    defer func() {
        requestDuration.WithLabelValues(r.URL.Path).Observe(time.Since(start).Seconds())
    }()
    // Handle request
})