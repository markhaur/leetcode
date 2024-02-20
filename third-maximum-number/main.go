package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Printf("third max element: %d\n", thirdMax(nums))

}

// https://leetcode.com/problems/third-maximum-number/
func thirdMax(nums []int) int {
	maxElems := []int{int(math.MinInt64), int(math.MinInt64), int(math.MinInt64)}

	for _, n := range nums {
		if n == maxElems[0] || n == maxElems[1] || n == maxElems[2] {
			continue
		}

		if n > maxElems[0] {
			if n > maxElems[2] {
				maxElems[2], maxElems[1], maxElems[0] = n, maxElems[2], maxElems[1]
				continue
			}
			if n > maxElems[1] {
				maxElems[1], maxElems[0] = n, maxElems[1]
				continue
			}
			maxElems[0] = n
		}
	}

	if maxElems[0] == int(math.MinInt64) {
		return maxElems[2]
	}

	return maxElems[0]
}
