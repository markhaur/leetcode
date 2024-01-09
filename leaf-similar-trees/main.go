package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := generateTree1()
	root2 := generateTree2()
	fmt.Printf("root1 and root2 has similar nodes: %t\n", leafSimilar(root1, root2))
}

// https://leetcode.com/problems/leaf-similar-trees/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(leaves(root1), leaves(root2))
}

func leaves(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	return append(leaves(root.Left), leaves(root.Right)...)
}

func generateTree1() *TreeNode {
	node6 := TreeNode{Val: 6}
	node7 := TreeNode{Val: 7}
	node4 := TreeNode{Val: 4}
	node9 := TreeNode{Val: 9}
	node8 := TreeNode{Val: 8}

	node2 := TreeNode{Val: 2, Left: &node7, Right: &node4}
	node5 := TreeNode{Val: 5, Left: &node6, Right: &node2}
	node1 := TreeNode{Val: 1, Left: &node9, Right: &node8}

	root := TreeNode{Val: 3, Left: &node5, Right: &node1}
	return &root
}

func generateTree2() *TreeNode {
	node6 := TreeNode{Val: 6}
	node7 := TreeNode{Val: 7}
	node4 := TreeNode{Val: 4}
	node9 := TreeNode{Val: 9}
	node8 := TreeNode{Val: 8}

	node2 := TreeNode{Val: 2, Left: &node9, Right: &node8}
	node5 := TreeNode{Val: 5, Left: &node6, Right: &node7}
	node1 := TreeNode{Val: 1, Left: &node4, Right: &node2}

	root := TreeNode{Val: 3, Left: &node5, Right: &node1}
	return &root
}
