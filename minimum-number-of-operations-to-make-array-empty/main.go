package main

import "fmt"

func main() {
	nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	fmt.Printf("min operations to delete the array: %d\n", minOperations(nums))
}

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
func minOperations(nums []int) int {
	store := make(map[int]int)

	for _, num := range nums {
		store[num]++
	}

	minOps := 0
	for _, val := range store {
		if val == 1 {
			return -1
		}

		if val%3 == 0 {
			minOps += val / 3
			continue
		}

		minOps += ((val / 3) + 1)
	}

	return minOps
}
