package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("Is Balanced Tree: %t\n", isBalanced(generateBalancedTree()))
	fmt.Printf("Is Balanced Tree: %t\n", isBalanced(nil))
	fmt.Printf("Is Balanced Tree: %t\n", isBalanced(generateUnbalancedTree()))
}

/**
* Given a binary tree, determine if it is height-balanced.
* A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
* @link: https://leetcode.com/problems/balanced-binary-tree/description/
 */
func isBalanced(root *TreeNode) bool {
	diff, _ := getDiff(root)
	return diff <= 1
}

// diff, depth
func getDiff(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftDiff, leftDepth := getDiff(node.Left)
	rightDiff, rightDepth := getDiff(node.Right)

	if leftDiff > 1 || rightDiff > 1 {
		return 2, 0
	}

	diff := leftDepth - rightDepth
	if diff < 0 {
		diff = -1 * diff
		return diff, rightDepth + 1
	}

	return diff, leftDepth + 1
}

func generateBalancedTree() *TreeNode {
	root := TreeNode{Val: 3}

	// first level
	l1First := TreeNode{Val: 9}
	l1Second := TreeNode{Val: 20}
	root.Left = &l1First
	root.Right = &l1Second

	// second level
	l2First := TreeNode{Val: 15}
	l2Second := TreeNode{Val: 7}
	l1Second.Left = &l2First
	l1Second.Right = &l2Second

	return &root
}

func generateUnbalancedTree() *TreeNode {
	root := TreeNode{Val: 1}

	// first level
	l1First := TreeNode{Val: 2}
	l1Second := TreeNode{Val: 2}
	root.Left = &l1First
	root.Right = &l1Second

	// second level
	l2First := TreeNode{Val: 3}
	l2Second := TreeNode{Val: 3}
	l1First.Left = &l2First
	l1Second.Right = &l2Second

	// third level
	l3First := TreeNode{Val: 4}
	l3Second := TreeNode{Val: 4}
	l2First.Left = &l3First
	l2Second.Right = &l3Second

	return &root
}
