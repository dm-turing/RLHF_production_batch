package main

import (
	"fmt"
	"sync"
	"time"
)

type LRUCache struct {
	capacity int
	cache    map[[2]string]*entry
	mu       sync.RWMutex
	head     *entry
	tail     *entry
}

type entry struct {
	key      [2]string
	value    interface{}
	prev     *entry
	next     *entry
	accessed time.Time
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		return nil
	}
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[[2]string]*entry),
	}
}
func (c *LRUCache) Get(namespace, key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	compositeKey := [2]string{namespace, key}
	if entry, ok := c.cache[compositeKey]; ok {
		c.moveToHead(entry)
		return entry.value
	}
	return nil
}

func (c *LRUCache) Put(namespace, key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	compositeKey := [2]string{namespace, key}
	if entry, ok := c.cache[compositeKey]; ok {
		entry.value = value
		c.moveToHead(entry)
		return
	}

	newEntry := &entry{
		key:      compositeKey,
		value:    value,
		accessed: time.Now(),
	}

	c.cache[compositeKey] = newEntry
	if len(c.cache) > c.capacity {
		c.removeTail()
	}

	c.addToHead(newEntry)
}

// Rest of the code remains the same...
func (c *LRUCache) moveToHead(entry *entry) {
	c.remove(entry)
	c.addToHead(entry)
}

func (c *LRUCache) addToHead(entry *entry) {
	entry.next = c.head
	entry.prev = nil
	if c.head != nil {
		c.head.prev = entry
	}
	c.head = entry
	if c.tail == nil {
		c.tail = entry
	}
}

func (c *LRUCache) remove(entry *entry) {
	if prev := entry.prev; prev != nil {
		prev.next = entry.next
	} else {
		c.head = entry.next
	}

	if next := entry.next; next != nil {
		next.prev = entry.prev
	} else {
		c.tail = entry.prev
	}
}

func (c *LRUCache) removeTail() {
	if c.tail == nil {
		return
	}
	c.remove(c.tail)
	delete(c.cache, c.tail.key)
	c.tail = c.tail.prev
}

func main() {
	cache := NewLRUCache(2)

	cache.Put("key1", "1", "value1")
	cache.Put("key2", "2", "value2")

	fmt.Println("Get key1:", cache.Get("key1", "1")) // Output: Get key1: value1
	fmt.Println("Get key3:", cache.Get("key3", "3")) // Output: Get key3: <nil>

	cache.Put("key3", "3", "value3")                 // key2 is evicted
	fmt.Println("Get key2:", cache.Get("key2", "2")) // Output: Get key2: <nil>
	fmt.Println("Get key3:", cache.Get("key3", "3")) // Output: Get key3: value3
}
