package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type FaultTolerantMap struct {
	data map[string]int
	mu   sync.Mutex
	file string
}

func NewFaultTolerantMap(file string) *FaultTolerantMap {
	ftm := &FaultTolerantMap{
		data: make(map[string]int),
		file: file,
	}
	// Load data from file if it exists
	if _, err := os.Stat(file); err == nil {
		ftm.loadData()
	}
	return ftm
}

func (ftm *FaultTolerantMap) Set(key string, value int) {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()

	ftm.data[key] = value
	ftm.saveData()
}

func (ftm *FaultTolerantMap) Get(key string) (int, bool) {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()

	return ftm.data[key], ftm.data[key] != 0
}

func (ftm *FaultTolerantMap) Delete(key string) {
	ftm.mu.Lock()
	defer ftm.mu.Unlock()

	delete(ftm.data, key)
	ftm.saveData()
}

func (ftm *FaultTolerantMap) saveData() {
	data, err := json.MarshalIndent(ftm.data, "", "  ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	err = ioutil.WriteFile(ftm.file, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (ftm *FaultTolerantMap) loadData() {
	data, err := ioutil.ReadFile(ftm.file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = json.Unmarshal(data, &ftm.data)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
		return
	}
}

func main() {
	ftm := NewFaultTolerantMap("faulttolerantmap.json")
	ftm.Set("key1", 1)
	ftm.Set("key2", 2)
	fmt.Println(ftm.Get("key1")) // Output: (1 true)
	ftm.Delete("key2")
	fmt.Println(ftm.Get("key2")) // Output: (0 false)
}
