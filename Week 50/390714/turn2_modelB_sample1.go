package main

import (
	"fmt"
)

// Simple decision tree node with a condition and actions
type TreeNode struct {
	Condition   func() bool
	TrueAction  func()
	FalseAction func()
}

func (t *TreeNode) Evaluate() {
	if t.Condition() {
		t.TrueAction()
	} else {
		t.FalseAction()
	}
}

// Global variables to represent the load balancer state
var serverCount int
var currentRequests int

// Decision tree functions
func isLowLoad() bool {
	return currentRequests <= 10
}

func scaleDown() {
	if serverCount > 1 {
		serverCount--
		fmt.Println("Scaling down to", serverCount, "servers.")
	} else {
		fmt.Println("Cannot scale down further.")
	}
}

func isHighLoad() bool {
	return currentRequests > 50
}

func scaleUp() {
	serverCount++
	fmt.Println("Scaling up to", serverCount, "servers.")
}

func maintainLoad() {
	fmt.Println("Current load is balanced.")
}

func runLoadBalancerDecisionTree() {
	root := &TreeNode{
		Condition:  isLowLoad,
		TrueAction: scaleDown,
		FalseAction: &TreeNode{
			Condition:   isHighLoad,
			TrueAction:  scaleUp,
			FalseAction: maintainLoad,
		}.Evaluate,
	}

	root.Evaluate()
}

func main() {
	// Simulate incoming requests
	currentRequests = 20 // Change this to test different scenarios
	runLoadBalancerDecisionTree()
}
