package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/ddgo"
	"github.com/DataDog/datadog-go/tracer"
)

func main() {

	// Initialize DataDog Tracer
	tracer.InitWithOptions(tracer.WithConfig(ddgo.Config{
		ClientToken: "YOUR_DATADOG_API_KEY", // Replace with your API key
		ServiceName: "order-service",
		Environment: "production",
	}))

	// Start a new trace for an incoming HTTP request
	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		span, ctx := tracer.StartSpanFromContext(r.Context(), "handleOrder", tracer.WithTag("method", r.Method), tracer.WithTag("endpoint", r.URL.Path))
		defer span.Finish()

		// Logic to handle the order
		processedOrder := processOrder(ctx)

		// Custom statistics
		tracer.Count("order.processed", 1, nil)
		tracer.Histogram("order.processing.time", float64(time.Now().Sub(span.StartTime().Add(time.Second)).Nanoseconds())*1e-6, nil)

		// Write the response
		fmt.Fprintf(w, "Processed Order: %s\n", processedOrder)
	})

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func processOrder(ctx context.Context) string {
	// Simulate processing time
	time.Sleep(2 * time.Second)

	// Start a new child span for the external call to ProcessService
	span, ctx := tracer.StartSpanFromContext(ctx, "sendOrderToProcessService", tracer.WithTag("service", "process-service"))
	defer span.Finish()

	// Simulate a request to ProcessService
	processedOrder := "Order Processed Successfully"

	return processedOrder
}
