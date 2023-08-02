package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Printf("%v intersection %v => %v\n", nums1, nums2, intersect(nums1, nums2))
}

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description
func intersect(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)

	for _, num1 := range nums1 {
		if numberExists(nums2, num1) {
			intersection = append(intersection, num1)
		}
	}
	return intersection
}

func numberExists(nums []int, num int) bool {
	for i, v := range nums {
		if v == num {
			nums[i] = -1
			return true
		}
	}
	return false
}
