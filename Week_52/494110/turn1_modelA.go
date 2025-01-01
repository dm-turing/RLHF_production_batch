package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{data: make(map[string]int)}
}

func (cm *ConcurrentMap) Set(key string, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.data[key] = value
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
	cm.mu.Lock() // Read lock
	defer cm.mu.Unlock()
	return cm.data[key], cm.data[key] != 0
}

func (cm *ConcurrentMap) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.data, key)
}

func main() {
	cm := NewConcurrentMap()
	go func() {
		for i := 0; i < 10000; i++ {
			cm.Set(fmt.Sprintf("key%d", i), i)
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			value, exists := cm.Get(fmt.Sprintf("key%d", i))
			if exists {
				fmt.Println(fmt.Sprintf("Get key%d: %d", i, value))
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
