import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("example.com/myapp")

func myFunction(ctx context.Context) {
	ctx, span := tracer.Start(ctx, "myFunction")
	defer span.End()

	// Add application logic here
}