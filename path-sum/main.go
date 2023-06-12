package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateBinaryTree()

	fmt.Printf("does binary tree container target sum: %d, %t\n", 22, hasPathSum(root, 22))
	fmt.Printf("does binary tree container target sum: %d, %t\n", 27, hasPathSum(root, 27))
	fmt.Printf("does binary tree container target sum: %d, %t\n", 19, hasPathSum(root, 19))
}

/**
* Given the `root` of a binary tree and an integer `targetSum`,
* return `true` if the tree has a root-to-leaf path such that adding up all the values along the path equals `targetSum`.
* A leaf is a node with no children.
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	return pathSum(root, 0, targetSum)
}

func pathSum(node *TreeNode, calcSum, targetSum int) bool {

	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return calcSum+node.Val == targetSum
	}

	hasSum := pathSum(node.Left, calcSum+node.Val, targetSum)
	if hasSum {
		return true
	}
	return pathSum(node.Right, calcSum+node.Val, targetSum)
}

/*
*                 5
*               4   8
*            11   13  4
*		   7   2        1
 */
func generateBinaryTree() *TreeNode {
	root := TreeNode{Val: 5}

	first1 := TreeNode{Val: 4}
	first2 := TreeNode{Val: 8}

	root.Left = &first1
	root.Right = &first2

	second1 := TreeNode{Val: 11}
	second2 := TreeNode{Val: 13}
	second3 := TreeNode{Val: 4}

	first1.Left = &second1
	first2.Left = &second2
	first2.Right = &second3

	third1 := TreeNode{Val: 7}
	third2 := TreeNode{Val: 2}
	third3 := TreeNode{Val: 1}

	second1.Left = &third1
	second1.Right = &third2
	second3.Right = &third3

	return &root

}
