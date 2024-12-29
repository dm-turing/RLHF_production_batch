package main

import (
	"fmt"
	"sync"
	"time"
)

type LRUCache struct {
	capacity int
	cache    map[compositeKey]*entry
	mu       sync.RWMutex
	head     *entry
	tail     *entry
}

type entry struct {
	key      compositeKey
	value    interface{}
	prev     *entry
	next     *entry
	accessed time.Time
}

type compositeKey struct {
	StringKey string
	IntKey    int
}

func (k compositeKey) Hash() uint32 {
	// Calculate a hash for the composite key
	return uint32(hashString(k.StringKey) ^ uint32(hashInt(k.IntKey)))
}

func (k compositeKey) Equal(other compositeKey) bool {
	// Compare the composite keys
	return k.StringKey == other.StringKey && k.IntKey == other.IntKey
}

func hashString(s string) uint32 {
	var h uint32 = 0
	for _, c := range s {
		h = 31*h + uint32(c)
	}
	return h
}

func hashInt(i int) uint32 {
	return uint32(i)
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		return nil
	}
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[compositeKey]*entry),
	}
}

func (c *LRUCache) Get(key compositeKey) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entry, ok := c.cache[key]; ok {
		c.moveToHead(entry)
		return entry.value
	}
	return nil
}

func (c *LRUCache) Put(key compositeKey, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.cache[key]; ok {
		entry.value = value
		c.moveToHead(entry)
		return
	}

	newEntry := &entry{
		key:      key,
		value:    value,
		accessed: time.Now(),
	}

	c.cache[key] = newEntry
	if len(c.cache) > c.capacity {
		c.removeTail()
	}

	c.addToHead(newEntry)
}

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

	cache.Put(compositeKey{"key1", 1}, "value1_1")
	cache.Put(compositeKey{"key2", 2}, "value2_2")

	fmt.Println("Get key1,1:", cache.Get(compositeKey{"key1", 1})) // Output: Get key1,1: value1_1
	fmt.Println("Get key3,3:", cache.Get(compositeKey{"key3", 3})) // Output: Get key3,3: <nil>

	cache.Put(compositeKey{"key3", 3}, "value3_3")                 // key2,2 is evicted
	fmt.Println("Get key2,2:", cache.Get(compositeKey{"key2", 2})) // Output: Get key2,2: <nil>
	fmt.Println("Get key3,3:", cache.Get(compositeKey{"key3", 3})) // Output: Get key3,3: value3_3
}
