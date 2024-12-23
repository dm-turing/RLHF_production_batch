package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	var cache sync.Map

	for n := 0; n < b.N; n++ {
		key := fmt.Sprintf("key-%d", n)
		value := fmt.Sprintf("Data for key %s", key)
		cache.Store(key, value)
		cache.Load(key)
	}
}
