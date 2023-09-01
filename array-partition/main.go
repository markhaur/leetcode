package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2, 6, 2, 6, 5, 1, 2}
	fmt.Printf("array pair sums: %d\n", arrayPairSum(nums))
}

// https://leetcode.com/problems/array-partition/description
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
		}
	}

	return sum
}
