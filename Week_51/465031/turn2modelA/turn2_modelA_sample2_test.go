package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkRegularMapMutex(b *testing.B) {
	var cache map[string]string = make(map[string]string)
	var mutex sync.Mutex

	for n := 0; n < b.N; n++ {
		key := fmt.Sprintf("key-%d", n)
		value := fmt.Sprintf("Data for key %s", key)

		mutex.Lock()
		cache[key] = value
		mutex.Unlock()

		mutex.Lock()
		_, _ = cache[key]
		mutex.Unlock()
	}
}
