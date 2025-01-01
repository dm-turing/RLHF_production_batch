package main

import (
	"sync"
)

var rw sync.RWMutex
var data = make(map[string]int)

func read(key string) int {
	rw.Lock()
	defer rw.Unlock()
	return data[key]
}

func write(key string, value int) {
	rw.Lock()
	defer rw.Unlock()
	data[key] = value
}

func main() {
	go func() { write("key", 10) }()
	go func() { read("key") }()
}
