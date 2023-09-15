package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Printf("findLHS: %d", findLHS(nums))
}

// https://leetcode.com/problems/longest-harmonious-subsequence/description
func findLHS(nums []int) int {
	store := make(map[int]int, 0)

	for _, num := range nums {
		store[num]++
	}

	maxSeq := 0
	count := 0
	numCount := 0
	ok := false

	for _, num := range nums {
		count, ok = store[num+1]
		if ok {
			numCount = store[num]
			if numCount+count > maxSeq {
				maxSeq = numCount + count
			}
		}
	}

	return maxSeq
}
