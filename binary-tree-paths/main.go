package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateBinaryTree()

	fmt.Printf("All Binary Tree Paths are: %v", binaryTreePaths(root))
}

func binaryTreePaths(root *TreeNode) []string {

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	paths := []string{}

	if root.Left != nil {
		paths = append(paths, binaryTreePaths(root.Left)...)
	}

	if root.Right != nil {
		paths = append(paths, binaryTreePaths(root.Right)...)
	}

	for i, path := range paths {
		paths[i] = fmt.Sprintf("%d->%s", root.Val, path)
	}

	return paths
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
