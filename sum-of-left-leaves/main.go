package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateBinaryTree()
	fmt.Printf("Sum of left leaves is: %d\n", sumOfLeftLeaves(root))
}

// https://leetcode.com/problems/sum-of-left-leaves/description
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum = root.Left.Val
	}

	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)

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
