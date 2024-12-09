package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	maxFailures   = 3
	retryInterval = 5 * time.Second
)

type CircuitBreaker struct {
	mu       sync.Mutex
	open     bool
	failures int
	lastFail time.Time
}

func (cb *CircuitBreaker) IsOpen() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.open
}

func (cb *CircuitBreaker) IncrementFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.failures++
	cb.lastFail = time.Now()
}

func (cb *CircuitBreaker) HalfOpen() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	if time.Now().Sub(cb.lastFail) > retryInterval {
		cb.open = false
		cb.failures = 0
	}
}

func processFile(filename string, cb *CircuitBreaker, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if cb.IsOpen() {
			log.Printf("Circuit breaker open for %s, retrying in %s", filename, retryInterval)
			time.Sleep(retryInterval)
			cb.HalfOpen()
			continue
		}

		if err := processFileInternal(filename, cb); err != nil {
			cb.IncrementFailure()
			if cb.failures >= maxFailures {
				cb.Open()
				log.Printf("Circuit breaker opened for %s", filename)
				return
			}
			continue
		}
		return
	}
}

func processFileInternal(filename string, cb *CircuitBreaker) error {
	// Simulate file processing. In a real application, this would be replaced with actual JSON or XML processing
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return err
	}

	select {
	case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
		log.Printf("Processed file: %s", filename)
		return nil
	default:
		log.Printf("Failed to process file: %s", filename)
		return fmt.Errorf("processing failed")
	}
}

func main() {
	start := time.Now()
	jsonFiles := []string{"user1.json", "user2.json"}
	xmlFiles := []string{"user1.xml", "user2.xml"}

	var wg sync.WaitGroup
	cb := &CircuitBreaker{}

	for _, file := range jsonFiles {
		wg.Add(1)
		go processFile(file, cb, &wg)
	}

	for _, file := range xmlFiles {
		wg.Add(1)
		go processFile(file, cb, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
