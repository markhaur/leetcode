package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Printf("search: %d", search(nums, target))
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	left := 0
	right := size - 1

	for left < right {
		midpoint := left + (right-left)/2
		if nums[midpoint] == target {
			return midpoint
		}
		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint
		}
	}

	start := left
	left = 0
	right = size - 1

	if target >= nums[start] && target <= nums[right] {
		left = start
	} else {
		right = start
	}

	for left <= right {
		midpoint := left + (right-left)/2
		if nums[midpoint] == target {
			return midpoint
		}
		if nums[midpoint] < target {
			left = midpoint + 1
		} else {
			right = midpoint - 1
		}
	}

	return -1
}
