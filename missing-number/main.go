package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Printf("missing number is %v is %d\n", missingNumber(nums), nums)
	nums = []int{0, 1}
	fmt.Printf("missing number is %v is %d\n", missingNumber(nums), nums)
	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Printf("missing number is %v is %d\n", missingNumber(nums), nums)
}

// https://leetcode.com/problems/missing-number/description
func missingNumber(nums []int) int {
	totalSum := (len(nums) * (2 + (len(nums) - 1))) / 2

	for _, num := range nums {
		totalSum -= num
	}

	return totalSum
}
