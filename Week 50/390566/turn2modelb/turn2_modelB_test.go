package turn2modelb

import (
	"testing"
	"time"
)

func TestTokenBucketBasic(t *testing.T) {
	tb := NewTokenBucket(5, 1)

	// Assert initial token count
	if tb.tokens != 5 {
		t.Errorf("Expected initial tokens to be 5, got %d", tb.tokens)
	}

	// Consume and verify token count
	if !tb.Consume() {
		t.Error("Failed to consume token")
	}
	if tb.tokens != 4 {
		t.Errorf("Expected token count to be 4, got %d", tb.tokens)
	}
}

func TestTokenBucketBursting(t *testing.T) {
	tb := NewTokenBucket(5, 1)
	// Consume all tokens initially
	for i := 0; i < 5; i++ {
		if !tb.Consume() {
			t.Errorf("Failed to consume token %d", i)
		}
	}

	// Wait for some time to refill tokens
	time.Sleep(time.Second)

	// Attempt to consume tokens beyond the burst
	for i := 0; i < 7; i++ {
		if i < 5 {
			if !tb.Consume() {
				t.Errorf("Failed to consume token %d during burst", i)
			}
		} else {
			if tb.Consume() {
				t.Errorf("Expected to be denied token %d during burst", i)
			}
		}
	}
}

// Add more test cases as per the aspects discussed above
