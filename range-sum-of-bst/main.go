package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateTree()
	low := 7
	high := 15
	fmt.Printf("range sum between %d and %d is %d\n", low, high, rangeSumBST(root, low, high))
}

// https://leetcode.com/problems/range-sum-of-bst/
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}

func generateTree() *TreeNode {
	node3 := TreeNode{Val: 3}
	node7 := TreeNode{Val: 7}
	node18 := TreeNode{Val: 18}

	node5 := TreeNode{Val: 5, Left: &node3, Right: &node7}
	node15 := TreeNode{Val: 15, Right: &node18}

	root := TreeNode{Val: 10, Left: &node5, Right: &node15}

	return &root
}
