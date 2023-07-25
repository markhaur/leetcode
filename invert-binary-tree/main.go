package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := generateBinaryTree()
	fmt.Printf("********* In-order traversal of Orignal Binary Tree *********\n")
	inorderTraversal(root)
	fmt.Println()

	root = invertTree(root)
	fmt.Printf("********* In-order traversal of Invert Binary Tree *********\n")
	inorderTraversal(root)
	fmt.Println()
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
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

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}
