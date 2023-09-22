package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("two sum: %v\n", twoSum(numbers, target))

}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {

	l, r, t := 0, len(numbers)-1, 0

	for {
		t = numbers[l] + numbers[r]
		if t == target {
			return []int{l + 1, r + 1}
		} else if t > target {
			r--
		} else if t < target {
			l++
		}
	}
}
