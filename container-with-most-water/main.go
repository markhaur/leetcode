package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("max area = %d\n", maxArea(height))
}

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	maxA := 0
	lowPtr := 0
	highPtr := len(height) - 1

	for lowPtr != highPtr {
		curA := min(height[lowPtr], height[highPtr]) * (highPtr - lowPtr)
		if curA > maxA {
			maxA = curA
		}

		if height[lowPtr] < height[highPtr] {
			lowPtr++
		} else {
			highPtr--
		}
	}

	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
