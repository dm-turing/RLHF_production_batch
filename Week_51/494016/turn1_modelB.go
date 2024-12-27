package main

import (
	"fmt"
	"sync"
)

type FaultTolerantMap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

// NewFaultTolerantMap creates a new FaultTolerantMap
func NewFaultTolerantMap() *FaultTolerantMap {
	return &FaultTolerantMap{m: make(map[string]interface{})}
}

// Put adds a key-value pair to the FaultTolerantMap
func (ftm *FaultTolerantMap) Put(key string, value interface{}) {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()
	ftm.m[key] = value
}

// Get retrieves a value for the specified key from the FaultTolerantMap
func (ftm *FaultTolerantMap) Get(key string) interface{} {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()
	return ftm.m[key]
}

// Delete deletes a key-value pair from the FaultTolerantMap
func (ftm *FaultTolerantMap) Delete(key string) {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()
	delete(ftm.m, key)
}

func main() {
	// Create a new fault-tolerant map
	ftm := NewFaultTolerantMap()

	// Insert some key-value pairs
	ftm.Put("name", "example")
	ftm.Put("age", 30)

	// Get value from the map
	name := ftm.Get("name").(string)
	fmt.Println("Name:", name)

	// Delete a key-value pair
	ftm.Delete("age")

	// Show the map state after deletion
	fmt.Println("Map state after deletion:", ftm.m)

	// Scenario 1:  Reading while another goroutine is writing (No data consistency)
	go func() {
		ftm.Put("location", "India")
	}()

	// Since the goroutine above is writing, this read could result in inconsistent data if the write operation is not synchronized.
	// we can use Locking mechanism to solve this issue
	location := ftm.Get("location").(string)
	fmt.Println("Location:", location)

	// Scenario 2:  Node failure or application crash (Data loss can occur)
	// Assuming no replication or persistence mechanism, if a node fails or the application crashes during writes or data updates,
	// some data could be lost.

}
