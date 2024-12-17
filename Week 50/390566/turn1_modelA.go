package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket implements the token bucket algorithm for rate limiting.
type TokenBucket struct {
	capacity int
	tokens   int
	fillRate int
	lastFill time.Time
	mutex    *sync.Mutex
}

// NewTokenBucket creates a new token bucket with given capacity and fill rate.
func NewTokenBucket(capacity, fillRate int) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		tokens:   capacity,
		fillRate: fillRate,
		lastFill: time.Now(),
		mutex:    &sync.Mutex{},
	}
}

// Consume tries to consume a token and return true if successful.
func (tb *TokenBucket) Consume() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	// Fill the bucket if needed
	tb.fill()

	if tb.tokens <= 0 {
		return false // No tokens available
	}

	tb.tokens--
	return true
}

// fill fills the token bucket with new tokens.
func (tb *TokenBucket) fill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastFill)
	secondsElapsed := elapsed.Seconds()
	tokensToAdd := int(secondsElapsed * float64(tb.fillRate))

	// Ensure tokens don't exceed the capacity
	if tokensToAdd > tb.capacity-tb.tokens {
		tokensToAdd = tb.capacity - tb.tokens
	}

	tb.tokens += tokensToAdd
	tb.lastFill = now
}

func main() {
	tb := NewTokenBucket(5, 1) // Allows 5 requests per second

	for i := 0; i < 20; i++ {
		if tb.Consume() {
			fmt.Println("Request", i, "allowed")
		} else {
			fmt.Println("Request", i, "denied")
			time.Sleep(time.Second) // Simulate retry or delay
		}
		time.Sleep(time.Millisecond * 500) // Simulate request processing time
	}
}
