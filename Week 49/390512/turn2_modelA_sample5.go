package main

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *List) Add(value int) {
	newNode := &Node{value: value}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *List) Get(index int) int {
	current := l.head
	for i := 0; i < index; i++ {
		if current == nil {
			panic(fmt.Sprintf("index %d out of bounds", index)) // Critical: accessing out of bounds
		}
		current = current.next
	}
	if current == nil {
		panic("unexpected nil node") // Critical: should not have a nil node here
	}
	return current.value
}

func main() {
	list := &List{}
	list.Add(1)
	list.Add(2)

	fmt.Println("Value at index 1:", list.Get(1)) // Works fine

	// This will panic
	fmt.Println("Value at index 3:", list.Get(3))
}
