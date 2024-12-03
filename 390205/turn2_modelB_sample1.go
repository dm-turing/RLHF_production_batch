package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type Stack struct {
	top *node
}

type node struct {
	value interface{}
	next  *node
}

func (s *Stack) Push(v interface{}) {
	n := &node{value: v}
	for {
		currentTop := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top))))
		n.next = currentTop
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)), unsafe.Pointer(currentTop), unsafe.Pointer(n)) {
			break
		}
	}
}

func (s *Stack) Pop() (interface{}, bool) {
	for {
		currentTop := (*node)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top))))
		if currentTop == nil {
			return nil, false
		}
		nextNode := currentTop.next
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&s.top)), unsafe.Pointer(currentTop), unsafe.Pointer(nextNode)) {
			return currentTop.value, true
		}
	}
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)

	value, ok := stack.Pop()
	if ok {
		fmt.Println("Popped:", value)
	}
}
