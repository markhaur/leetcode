package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 8, 1, 5, 4, 5, 2, 10, 3, 6, 5, 2, 3}
	fmt.Printf("max product: %d\n", maxProduct(nums))
}

// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array
func maxProduct(nums []int) int {
	maxNums := []int{nums[0], nums[1]}

	for i := 2; i < len(nums); i++ {
		if nums[i] > maxNums[0] {
			if maxNums[0] > maxNums[1] {
				maxNums[1] = maxNums[0]
			}
			maxNums[0] = nums[i]
			continue
		}
		if nums[i] > maxNums[1] {
			maxNums[1] = nums[i]
		}
	}

	return (maxNums[0] - 1) * (maxNums[1] - 1)
}
