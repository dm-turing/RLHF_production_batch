package main

import (
	"sync"
	"testing"
)

const numGoroutines = 100
const numIncrements = 1000

// Using Mutex
type MutexCounter struct {
	mu    sync.Mutex
	count int
}

func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func MutexBenchmark(b *testing.B) {
	var counter MutexCounter
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			go func() {
				for k := 0; k < numIncrements; k++ {
					counter.Increment()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// Using Channels
type ChannelCounter struct {
	count chan int
}

func NewChannelCounter() *ChannelCounter {
	cc := &ChannelCounter{
		count: make(chan int),
	}

	go func() {
		total := 0
		for increment := range cc.count {
			total += increment
		}
		cc.count <- total // Send final count back
	}()
	return cc
}

func (cc *ChannelCounter) Increment() {
	cc.count <- 1
}

func ChannelBenchmark(b *testing.B) {
	cc := NewChannelCounter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			go func() {
				for k := 0; k < numIncrements; k++ {
					cc.Increment()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(cc.count) // Close the channel after all increments
		<-cc.count      // Wait for the final count
	}
}

func main() {
	// This is just to show how you would typically run the benchmarks.
	// Use "go test -bench=." in the terminal to execute these benchmarks.
}
