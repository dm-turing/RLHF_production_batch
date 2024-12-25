package main

import (
	"sync"
	"testing"
)

var (
	sharedSlice = make([]int, 100000)
	mutex       sync.RWMutex
)

func BenchmarkCopyConcurrentByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyConcurrentByValue()
	}
}

func copyConcurrentByValue() {
	mutex.RLock()
	defer mutex.RUnlock()
	copyByValue(sharedSlice)
}

func BenchmarkCopyConcurrentByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyConcurrentByReference()
	}
}

func copyConcurrentByReference() {
	mutex.RLock()
	defer mutex.RUnlock()
	copyByReference(&sharedSlice)
}

func BenchmarkAppendConcurrentByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendConcurrentByValue()
	}
}

func appendConcurrentByValue() {
	mutex.Lock()
	defer mutex.Unlock()
	appendByValue(sharedSlice)
}

func BenchmarkAppendConcurrentByReference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendConcurrentByReference()
	}
}

func appendConcurrentByReference() {
	mutex.Lock()
	defer mutex.Unlock()
	appendByReference(&sharedSlice)
}
