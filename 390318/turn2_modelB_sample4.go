package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("example.com/myapp")

func myFunction(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "myFunction")
	defer span.End()

	// Add attributes to the span
	span.SetAttributes(attribute.String("key", "value"))

	// Application logic
}

func main() {
	ctx := context.Background()
	myFunction(ctx)
}
