package main

import (
	"fmt"
	"sync"
)

// TaskGraph structure to model dependencies
type TaskGraph struct {
	tasks map[string][]string
	mu    sync.Mutex
}

// NewTaskGraph initializes a new TaskGraph
func NewTaskGraph() *TaskGraph {
	return &TaskGraph{
		tasks: make(map[string][]string),
	}
}

// AddDependency adds a dependency between tasks
func (tg *TaskGraph) AddDependency(task, dependency string) {
	tg.mu.Lock()
	defer tg.mu.Unlock()
	tg.tasks[task] = append(tg.tasks[task], dependency)
}

// TopologicalSort performs a topological sort on the task graph
func (tg *TaskGraph) TopologicalSort() []string {
	tg.mu.Lock()
	defer tg.mu.Unlock()

	visited := make(map[string]bool)
	stack := []string{}
	var visitAll func(tasks []string)

	visitAll = func(tasks []string) {
		for _, task := range tasks {
			if !visited[task] {
				visited[task] = true
				visitAll(tg.tasks[task])
				stack = append(stack, task)
			}
		}
	}

	for task := range tg.tasks {
		visitAll([]string{task})
	}

	// Reverse the stack to get the order of execution
	for i := len(stack)/2 - 1; i >= 0; i-- {
		opp := len(stack) - 1 - i
		stack[i], stack[opp] = stack[opp], stack[i]
	}

	return stack
}

func main() {
	// Create a new task graph
	taskGraph := NewTaskGraph()

	// Add dependencies
	taskGraph.AddDependency("task3", "task1")
	taskGraph.AddDependency("task3", "task2")
	taskGraph.AddDependency("task2", "task1")

	// Get topological order
	order := taskGraph.TopologicalSort()
	fmt.Println("Task execution order:", order)
}
