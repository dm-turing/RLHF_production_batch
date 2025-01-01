package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentMap struct {
	m sync.Map
}

func (cm *ConcurrentMap) Store(key, value interface{}) {
	cm.m.Store(key, value)
}

func (cm *ConcurrentMap) Load(key interface{}) (interface{}, bool) {
	return cm.m.Load(key)
}

func (cm *ConcurrentMap) Delete(key interface{}) {
	cm.m.Delete(key)
}

func (cm *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cm.m.Range(f)
}

func main() {
	// Create a new concurrent map
	cm := &ConcurrentMap{}

	// Use goroutines to add/remove key-value pairs concurrently
	go func() {
		cm.Store("key1", 100)
		cm.Store("key2", 200)
		cm.Delete("key1")
	}()

	go func() {
		cm.Store("key3", 300)
		cm.Store("key4", 400)
		cm.Delete("key3")
	}()

	// Wait for the goroutines to finish their work
	time.Sleep(time.Second)

	// Print the final values in the map
	fmt.Println("Final Map Values:")
	cm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // continue the range
	})
}
