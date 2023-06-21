package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("min depth of normal tree is %d\n", minDepth(generateTree()))
	fmt.Printf("min depth of right tree is %d\n", minDepth(generateRightTree()))
}

/**
* Given a binary tree, find its minimum depth.
* The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
* Note: A leaf is a node with no children.
* @link: https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if leftDepth == 0 {
		return rightDepth + 1
	}

	if rightDepth == 0 {
		return leftDepth + 1
	}

	if leftDepth < rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func generateTree() *TreeNode {
	root := TreeNode{Val: -9}

	// level 1 nodes
	l1First := TreeNode{Val: -3}
	l1Second := TreeNode{Val: 2}
	root.Left = &l1First
	root.Right = &l1Second

	// level 2 nodes
	l2First := TreeNode{Val: 4}
	l2Second := TreeNode{Val: 4}
	l2Third := TreeNode{Val: 0}

	l1First.Right = &l2First
	l1Second.Left = &l2Second
	l1Second.Right = &l2Third

	// level 3 nodes
	l3First := TreeNode{Val: -6}
	l3Second := TreeNode{Val: -5}

	l2First.Left = &l3First
	l2Second.Right = &l3Second

	return &root
}

func generateRightTree() *TreeNode {
	root := TreeNode{Val: 2}

	first := TreeNode{Val: 3}
	root.Right = &first

	second := TreeNode{Val: 4}
	first.Right = &second

	third := TreeNode{Val: 5}
	second.Right = &third

	fourth := TreeNode{Val: 6}
	third.Right = &fourth

	return &root
}
