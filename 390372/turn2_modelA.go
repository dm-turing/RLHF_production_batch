package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

const (
    maxFailures  = 3
    successThreshold = 3
    checkInterval     = 5 * time.Second
)

type CircuitBreaker struct {
    mu         sync.Mutex
    failures    int32
    successes   int32
    isOpen      bool
    lastCheck   time.Time
    successRate float32
}

func (cb *CircuitBreaker) IsOpen() bool {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    return cb.isOpen
}

func (cb *CircuitBreaker) MarkSuccess() {
    cb.mu.Lock()
    defer cb.mu.Unlock()

    atomic.AddInt32(&cb.successes, 1)
    cb.updateSuccessRate()

    if cb.isOpen {
        cb.halfOpen()
    }
}

func (cb *CircuitBreaker) MarkFailure() {
    cb.mu.Lock()
    defer cb.mu.Unlock()

    atomic.AddInt32(&cb.failures, 1)
    cb.updateSuccessRate()

    if cb.failures >= maxFailures {
        cb.trip()
    }
}

func (cb *CircuitBreaker) trip() {
    cb.mu.Lock()
    defer cb.mu.Unlock()

    cb.failures = 0
    cb.isOpen = true
    go cb.checkHealth()
}

func (cb *CircuitBreaker) checkHealth() {
    for {
        time.Sleep(checkInterval)
        cb.mu.Lock()
        defer cb.mu.Unlock()

        if !cb.isOpen {
            break
        }

        if cb.successes >= successThreshold {
            cb.close()
        }
    }
}

func (cb *CircuitBreaker) close() {
    cb.mu.Lock()
    defer cb.mu.Unlock()

    cb.isOpen = false
    cb.successes = 0
}

func (cb *CircuitBreaker) updateSuccessRate() {
    total := cb.successes + cb.failures
    if total == 0 {
        cb.successRate = 0.0
        return
    }
    cb.successRate = float32(cb.successes) / float32(total)
}

func processFile(filename string, cb *CircuitBreaker) {
    if cb.IsOpen() {
        fmt.Println("Circuit breaker is open. Skipping file:", filename)
        return
    }

    // Simulate processing the file
    if err := simulateProcessing(filename); err != nil {
        cb.MarkFailure()
    } else {
        cb.MarkSuccess()
    }
}

func simulateProcessing(filename string) error {
    // Randomly decide if the processing fails or succeeds
    if time.Now().UnixNano()%2 == 0 {
        return fmt.Errorf("Failed to process file: %s", filename)
    }
    return nil
}

func main() {
    files := []string{"file1.json", "file2.json", "file3.xml", "file4.xml", "file5.json"}
    cb := &CircuitBreaker{}