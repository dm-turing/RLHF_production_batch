package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket struct represents the token bucket
type TokenBucket struct {
	capacity int        // maximum number of tokens
	rate     int        // rate of token generation (tokens per second)
	tokens   int        // current number of tokens
	lastFill time.Time  // last time the bucket was filled
	lock     sync.Mutex // to ensure thread safety
}

// NewTokenBucket creates a new token bucket with the given capacity and rate
func NewTokenBucket(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		rate:     rate,
		tokens:   capacity,
		lastFill: time.Now(),
	}
}

// IsFull checks if the token bucket is full
func (tb *TokenBucket) IsFull() bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()
	return tb.tokens >= tb.capacity
}

// Take takes a token from the bucket if available
func (tb *TokenBucket) Take() bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()

	// If the bucket is full, return false
	if tb.IsFull() {
		return false
	}

	// Calculate the number of tokens that can be added since the last fill
	now := time.Now()
	delta := now.Sub(tb.lastFill)
	numTokensToAdd := int(delta.Seconds() * float64(tb.rate))

	// Add the tokens and limit to the capacity
	tb.tokens += numTokensToAdd
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	// If there are still tokens available, take one and return true
	if tb.tokens > 0 {
		tb.tokens--
		tb.lastFill = now
		return true
	}

	return false
}

func main() {
	// Create a new token bucket with a capacity of 10 tokens and a rate of 1 token per second
	bucket := NewTokenBucket(5, 1)

	for i := 0; i < 20; i++ {
		if bucket.Take() {
			// Request is allowed
			fmt.Printf("Request %d is allowed.\n", i+1)
			time.Sleep(time.Second)
		} else {
			// Request is rate limited
			fmt.Printf("Request %d is rate limited.\n", i+1)
		}
	}
}
