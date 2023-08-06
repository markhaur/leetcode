package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("Binary Tree Total Nodes: %d", countNodes(generateBinaryTree()))
}

// https://leetcode.com/problems/count-complete-tree-nodes/description
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func generateBinaryTree() *TreeNode {
	root := TreeNode{Val: 4}

	l1First := TreeNode{Val: 2}
	l1Second := TreeNode{Val: 7}
	root.Left = &l1First
	root.Right = &l1Second

	l2First := TreeNode{Val: 1}
	l2Second := TreeNode{Val: 3}
	l1First.Left = &l2First
	l1First.Right = &l2Second

	l2Third := TreeNode{Val: 6}
	l2Fourth := TreeNode{Val: 9}
	l1Second.Left = &l2Third
	l1Second.Left = &l2Fourth

	return &root
}
