package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	fmt.Printf("max consecutive ones: %d\n", findMaxConsecutiveOnes(nums))
}

// https://leetcode.com/problems/max-consecutive-ones/description
func findMaxConsecutiveOnes(nums []int) int {
	maxOne, count := 0, 0

	for _, n := range nums {
		if n == 0 {
			if count > maxOne {
				maxOne = count
			}
			count = -1
		}
		count++
	}

	if count > maxOne {
		maxOne = count
	}

	return maxOne
}
