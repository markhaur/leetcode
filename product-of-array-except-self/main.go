package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("product except self: %v\n", productExceptSelf(nums))
}

// https://leetcode.com/problems/product-of-array-except-self/description
func productExceptSelf(nums []int) []int {
	product := 1
	isZero := false
	for _, n := range nums {
		if isZero && n == 0 {
			product = 0
			break
		}
		if n == 0 {
			isZero = true
			continue
		}
		product *= n
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = product
		} else if isZero {
			nums[i] = 0
		} else {
			nums[i] = product / nums[i]
		}
	}

	return nums
}
