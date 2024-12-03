package main

import (
	"fmt"
	"sync/atomic"
)

var counter uint64 = 0

func incrementCounter() {
	atomic.AddUint64(&counter, 1)
}

var value uint64 = 42

func getValue() uint64 {
	return atomic.LoadUint64(&value)
}

func setValue(newValue uint64) {
	atomic.StoreUint64(&value, newValue)
}

func swapValue(newValue uint64) uint64 {
	return atomic.SwapUint64(&value, newValue)
}

func compareAndSwap(expected uint64, newValue uint64) bool {
	return atomic.CompareAndSwapUint64(&value, expected, newValue)
}

var intValue int32 = 0

func swapInt(newValue int32) {
	atomic.SwapInt32(&intValue, newValue)
}

var int64Value int64 = 100

func loadInt64() int64 {
	return atomic.LoadInt64(&int64Value)
}

func storeInt64(newValue int64) {
	atomic.StoreInt64(&int64Value, newValue)
}

func main() {
	incrementCounter()
	fmt.Println(counter)
	getValue()
	fmt.Println(value)
	setValue(1000)
	fmt.Println(value)
	swapValue(20000)
	fmt.Println(value)
	compareAndSwap(atomic.LoadUint64(&value), 30000)
	fmt.Println(value)
	swapInt(20000)
	fmt.Println(intValue)
	fmt.Println(loadInt64())
	storeInt64(30000)
	fmt.Println(int64Value)
}
