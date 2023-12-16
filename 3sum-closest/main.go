package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Printf("three sum closest: %d\n", threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	var ans int
	var imin int
	var imax int
	var difftemp int
	var sum int

	diff := 100000
	sort.Ints(nums)
	leng := len(nums)

	for index := 0; index < leng-2; index++ {
		imin = index + 1
		imax = leng - 1
		for imin < imax {
			sum = nums[index] + nums[imin] + nums[imax]
			difftemp = sum - target

			//make the diff aways positive
			if difftemp < 0 {
				difftemp = -difftemp
			}

			//check if it is the smallest diff
			if difftemp < diff {
				diff = difftemp
				ans = sum
			}

			if sum == target {
				return target // best case

			} else if sum > target {
				imax-- // make the sum smaller
			} else {
				imin++ // make the sum bigger
			}
		}
	}
	return ans
}
