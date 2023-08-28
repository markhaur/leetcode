package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Printf("next greater elements: %v\n", nextGreaterElement(nums1, nums2))
}

// https://leetcode.com/problems/next-greater-element-i/description
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums1Length := len(nums1)
	nums2Length := len(nums2)

	store := make(map[int]int, nums2Length)
	gElem := -1
	for i, _ := range nums2 {
		gElem = -1
		for z := i + 1; z < nums2Length; z++ {
			if nums2[z] > nums2[i] {
				gElem = nums2[z]
				break
			}
		}
		store[nums2[i]] = gElem
	}

	result := make([]int, nums1Length)
	for i, n1 := range nums1 {
		result[i] = store[n1]
	}

	return result
}
