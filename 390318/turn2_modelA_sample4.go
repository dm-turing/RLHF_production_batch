package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/http"
)

var requestsCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name:       "request_count_total",
		Help:       "Total number of requests received.",
		Subsystem:  "http",
		Registerer: prometheus.DefaultRegisterer,
	},
	[]string{"method", "code"},
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestsCount.WithLabelValues(r.Method, strconv.Itoa(http.StatusOK)).Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/health", requestHandler)
	http.Handle("/metrics", prometheus.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
