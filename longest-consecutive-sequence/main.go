package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Printf("longest consecutive: %d\n", longestConsecutive(nums))

}

// https://leetcode.com/problems/longest-consecutive-sequence/description
func longestConsecutive(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	sort.Ints(nums)
	lCons := 0
	cons := 1

	for i := 0; i < size-1; i++ {
		if nums[i] != nums[i+1] {
			if nums[i]+1 == nums[i+1] {
				cons++
			} else {
				if cons > lCons {
					lCons = cons
				}
				cons = 1
			}
		}
	}

	if cons > lCons {
		return cons
	}

	return lCons
}
