package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Printf("tree 2 str: %s\n", tree2str(generateTree()))
}

// https://leetcode.com/problems/construct-string-from-binary-tree
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := fmt.Sprintf("%d", root.Val)
	if root.Left != nil {
		res += fmt.Sprintf("(%s)", tree2str(root.Left))
	} else if root.Right != nil {
		res += "()"
	}
	if root.Right != nil {
		res += fmt.Sprintf("(%s)", tree2str(root.Right))
	}
	return res
}

func generateTree() *TreeNode {
	node4 := TreeNode{Val: 4}
	node3 := TreeNode{Val: 3}
	node2 := TreeNode{Val: 2, Left: &node4}
	root := TreeNode{Val: 1, Left: &node2, Right: &node3}

	return &root
}
