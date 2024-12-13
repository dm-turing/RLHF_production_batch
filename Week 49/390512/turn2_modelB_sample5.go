package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func isBST(node *TreeNode) bool {
	return isBSTHelper(node, nil, nil)
}

func isBSTHelper(node *TreeNode, minValue *int, maxValue *int) bool {
	if node == nil {
		return true
	}

	if minValue != nil && *minValue >= node.Value {
		panic("binary search tree invariant violated: node value less than or equal to minValue")
	}

	if maxValue != nil && *maxValue <= node.Value {
		panic("binary search tree invariant violated: node value greater than or equal to maxValue")
	}

	return isBSTHelper(node.Left, minValue, &node.Value) && isBSTHelper(node.Right, &node.Value, maxValue)
}

func main() {
	root := &TreeNode{Value: 2}
	root.Left = &TreeNode{Value: 1}
	root.Right = &TreeNode{Value: 3}

	root.Left.Right = &TreeNode{Value: 1} // Invalid for BST

	if !isBST(root) {
		fmt.Println("Tree is not a BST.")
	}
}
