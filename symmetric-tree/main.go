package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateTree()
	fmt.Printf("is tree symmetric: %t\n", isSymmetric(root))
}

/**
 * Given the `root` of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
 * https://leetcode.com/problems/symmetric-tree/description/
 */
func isSymmetric(root *TreeNode) bool {
	return symmetricSubtree(root.Left, root.Right)
}

func symmetricSubtree(leftNode, rightNode *TreeNode) bool {

	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil {
		return false
	}

	if leftNode.Val != rightNode.Val {
		return false
	}

	leftSymmetric := symmetricSubtree(leftNode.Left, rightNode.Right)
	rightSymmetric := symmetricSubtree(leftNode.Right, rightNode.Left)

	return leftSymmetric && rightSymmetric
}

func generateTree() *TreeNode {
	root := TreeNode{Val: 1}

	firstNode := TreeNode{Val: 2}
	secondNode := TreeNode{Val: 2}
	root.Left = &firstNode
	root.Right = &secondNode

	fourthNode := TreeNode{Val: 3}
	fifthNode := TreeNode{Val: 4}
	firstNode.Left = &fourthNode
	firstNode.Right = &fifthNode

	sixthNode := TreeNode{Val: 4}
	seventhNode := TreeNode{Val: 3}
	secondNode.Left = &sixthNode
	secondNode.Right = &seventhNode

	return &root
}
