package main

import (
	"sync"
)

var sharedData = &struct{}{}

func criticalSection(mutex *sync.Mutex) {
	// Assume this section contains critical resources
	mutex.Lock()
	defer mutex.Unlock()
	if sharedData == nil {
		panic("sharedData is nil, application state inconsistent")
	}
	// Critical operations
}

func main() {
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		criticalSection(mutex)
	}()

	go func() {
		defer wg.Done()
		sharedData = nil       // Simulating an issue
		criticalSection(mutex) // Will panic
	}()

	wg.Wait()
}
