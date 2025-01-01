package main

import "sync"

type ConcurrentMap1 struct {
	m   sync.Map
	rwl sync.RWMutex
}

func (cm *ConcurrentMap1) Lock(key interface{}) {
	cm.rwl.RLock()
	_, loaded := cm.m.Load(key)
	cm.rwl.RUnlock()

	// Use exclusive lock for writes
	cm.rwl.Lock()
	if !loaded {
		cm.m.Store(key, struct{}{})
	}
	cm.rwl.Unlock()
}

func (cm *ConcurrentMap1) Unlock(key interface{}) {
	cm.rwl.Lock()
	cm.m.Delete(key)
	cm.rwl.Unlock()
}

func main() {
	c := &ConcurrentMap1{}
	c.Lock("key")
	c.Unlock("key")
}
