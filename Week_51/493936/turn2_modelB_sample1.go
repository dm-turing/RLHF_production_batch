package main

import (
	"context"
	"fmt"
)

type ContextError interface {
	error
	Context() context.Context
}

type OrderError struct {
	Ctx     context.Context
	Message string
}

func (e *OrderError) Error() string {
	return fmt.Sprintf("order error: %s (at: %v)", e.Message, e.Ctx.Value("location"))
}

func (e *OrderError) Context() context.Context {
	return e.Ctx
}

func processOrder(ctx context.Context, orderID int) error {
	if orderID == 0 {
		return &OrderError{
			Ctx:     ctx,
			Message: "order ID cannot be zero",
		}
	}
	// ... process the order
	return nil
}

func main() {
	ctx := context.WithValue(context.Background(), "location", "main.go:22")
	err := processOrder(ctx, 0)
	if cerr, ok := err.(ContextError); ok {
		fmt.Printf("Error context: %v\n", cerr.Context().Value("location"))
	}
	fmt.Println(err)
}
