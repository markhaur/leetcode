package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("Minimum Value of BST: %d\n", getMinimumDifference(generateBST()))
}

/**
*   	4
*     2   6
*   1  3
 */
func generateBST() *TreeNode {
	root := TreeNode{Val: 4}

	first1 := TreeNode{Val: 2}
	first2 := TreeNode{Val: 6}

	root.Left = &first1
	root.Right = &first2

	second1 := TreeNode{Val: 1}
	second2 := TreeNode{Val: 3}

	first1.Left = &second1
	first1.Right = &second2

	return &root
}

// Given the `root` of a Binary Search Tree (BST),
// return the minimum absolute difference between the values of any two different nodes in the tree.
func getMinimumDifference(root *TreeNode) int {
	min := 100001
	prev := 100001
	minDifference(root, &min, &prev)
	return min
}

func minDifference(node *TreeNode, min *int, prev *int) {

	if node == nil {
		return
	}

	minDifference(node.Left, min, prev)

	*prev = *prev - node.Val
	if *prev < 0 {
		*prev = *prev * -1
	}
	if *prev < *min {
		*min = *prev
	}
	*prev = node.Val

	minDifference(node.Right, min, prev)
}
