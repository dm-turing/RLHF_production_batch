package main

import (
	"fmt"
	"sync"
)

// Cache represents a simple caching system using a map
type Cache struct {
	cache map[string]interface{}
	mutex *sync.RWMutex
}

// NewCache creates a new instance of the Cache
func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
		mutex: &sync.RWMutex{},
	}
}

// Set adds a key-value pair to the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
}

// Get retrieves a value from the cache based on the key
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.cache[key]
	return value, ok
}

// Invalidate removes a key-value pair from the cache
func (c *Cache) Invalidate(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, key)
}

// FullInvalidate clears the entire cache
func (c *Cache) FullInvalidate() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache = make(map[string]interface{})
}

// Let's test the cache implementation
func main() {
	cache := NewCache()

	// Cache miss (no data in the cache for "city")
	value, ok := cache.Get("city")
	if !ok {
		fmt.Println("Cache miss: Key not found")
	}
	// Set a new key-value pair
	cache.Set("city", "San Francisco")

	// Cache hit (data exists in the cache for "city")
	value, ok = cache.Get("city")
	if ok {
		fmt.Println("Cache hit:", value)
	}

	// Cache miss after invalidation
	cache.Invalidate("city")
	value, ok = cache.Get("city")
	if !ok {
		fmt.Println("Cache miss after invalidation: Key not found")
	}
}
