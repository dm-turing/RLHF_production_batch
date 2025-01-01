package main

import (
	"hash/maphash"
	"sync"
)

const numStripes = 100

var stripes = make([]sync.Mutex, numStripes)

func hashKey(key string) int {
	h := maphash.Hash{}
	h.SetSeed(maphash.MakeSeed())
	h.WriteString(key)
	return int(h.Sum64() % uint64(numStripes))
}

func store(key, value string) {
	idx := hashKey(key)
	stripes[idx].Lock()
	defer stripes[idx].Unlock()
	// Store the value
}

func load(key string) string {
	idx := hashKey(key)
	stripes[idx].Lock()
	defer stripes[idx].Unlock()
	// Load the value
	return ""
}

func main() {
	go func() { store("key1", "value1") }()
	go func() { store("key2", "value2") }()
	go func() { load("key1") }()
	go func() { load("key2") }()
}
