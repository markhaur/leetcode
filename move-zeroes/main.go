package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 3, 4, 5, 6, 7, 8, 9}
	moveZeroes(nums)
	fmt.Printf("after moving zeroes: %v\n", nums)
}

// https://leetcode.com/problems/move-zeroes/description
func moveZeroes(nums []int) {
	swapDistance := 0
	for i, num := range nums {
		if num == 0 {
			swapDistance += 1
		} else if swapDistance > 0 {
			nums[i-swapDistance] = num
			nums[i] = 0
		}
	}
}
