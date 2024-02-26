package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := generateTree()
	fmt.Printf("is same tree: %t\n", isSameTree(tree, tree))
}

// https://leetcode.com/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func generateTree() *TreeNode {
	two := TreeNode{Val: 2}
	three := TreeNode{Val: 3}

	root := TreeNode{Val: 1, Left: &two, Right: &three}
	return &root
}
