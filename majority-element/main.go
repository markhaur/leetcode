package main

import "fmt"

func main() {
	fmt.Printf("majority element is %d in [3,2,3]\n", majorityElement([]int{3, 2, 3}))
	fmt.Printf("majority element is %d in [2,2,1,1,1,2,2]\n", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Printf("majority element is %d in [3,2,3,4,5,6,7,5,5,5,5,5,5,5]\n", majorityElement([]int{3, 2, 3, 4, 5, 6, 7, 5, 5, 5, 5, 5, 5, 5}))
}

// Given an array `nums` of size `n`, return the majority element.
// The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	countMem := make(map[int]int)
	maxCount := 0
	maxNum := -1

	for _, num := range nums {
		count := countMem[num]
		count = count + 1
		countMem[num] = count
		if count > maxCount {
			maxNum = num
			maxCount = count
		}
	}

	return maxNum
}
