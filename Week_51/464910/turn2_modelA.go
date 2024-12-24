package main

import (
	"fmt"
	"time"
)

// Callback types for different stages
type orderCreatedCallback func()
type orderProcessingStartedCallback func()
type orderShippedCallback func()
type orderCompletedCallback func()

// Functions to represent stages of order processing
func createOrder(createdCallback orderCreatedCallback) {
	fmt.Println("Order created!")
	time.Sleep(100 * time.Millisecond) // Simulate delay
	createdCallback()
}

func processOrder(orderId int, startCallback orderProcessingStartedCallback, shippingCallback orderShippedCallback, completeCallback orderCompletedCallback) {
	fmt.Println("Order processing started...")
	time.Sleep(500 * time.Millisecond) // Simulate delay
	startCallback()

	time.Sleep(300 * time.Millisecond) // Simulate delay
	fmt.Printf("Shipping order %d...\n", orderId)
	shippingCallback()

	time.Sleep(400 * time.Millisecond) // Simulate delay
	fmt.Println("Order completed.")
	completeCallback()
}

// Stage handler triggering necessary callbacks
func handleOrder(orderId int) {
	createdCallback := func() {
		fmt.Println("Notifying creation completion...")
		// Register callbacks for the processing stage
		registerProcessingCallbacks(orderId)
	}

	processStartCallback := func() {
		fmt.Println("Notifying processing started...")
	}

	orderShippedCallback := func() {
		fmt.Println("Notifying order shipped...")
	}

	orderCompleteCallback := func() {
		fmt.Println("Notifying order completed...")
	}

	// Trigger initial stage
	createOrder(createdCallback)
	processStartCallback()
	orderShippedCallback()
	orderCompleteCallback()
}

func registerProcessingCallbacks(orderId int) {
	// Declare callbacks
	startCallback := func() {
		fmt.Println("Processing started callback triggered.")
	}

	shipCallback := func() {
		fmt.Println("Shipping callback triggered.")
	}

	completeCallback := func() {
		fmt.Println("Completion callback triggered.")
	}

	// Trigger processing stage
	processOrder(orderId, startCallback, shipCallback, completeCallback)
}

// Main function to start the event management system
func main() {
	handleOrder(12345)
}
