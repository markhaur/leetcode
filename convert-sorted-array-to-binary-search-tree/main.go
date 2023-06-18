package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// Given an integer array `nums` where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
func sortedArrayToBST(nums []int) *TreeNode {
	leftRange := 0
	rightRange := len(nums) - 1
	return build(nums, leftRange, rightRange)
}

func build(nums []int, leftRange, rightRange int) *TreeNode {
	if leftRange > rightRange {
		return nil
	}

	mid := (leftRange + rightRange) / 2

	root := TreeNode{Val: nums[mid]}
	root.Left = build(nums, leftRange, mid-1)
	root.Right = build(nums, mid+1, rightRange)
	return &root
}
