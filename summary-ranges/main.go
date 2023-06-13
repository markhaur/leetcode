package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Printf("ranges of `nums` array: %v", summaryRanges(nums))
}

/**
You are given a sorted unique integer array nums.
A range [a,b] is the set of all integers from a to b (inclusive).
Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
Each range [a,b] in the list should be output as:
"a->b" if a != b
"a" if a == b
*/
func summaryRanges(nums []int) []string {
	var ranges []string
	startIndex := 0
	size := len(nums)

	for i := 0; i <= size-1; i++ {

		if i+1 == size || nums[i]+1 != nums[i+1] {
			if startIndex == i {
				ranges = append(ranges, fmt.Sprintf("%d", nums[startIndex]))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d->%d", nums[startIndex], nums[i]))
			}
			startIndex = i + 1
		}
	}

	return ranges
}
