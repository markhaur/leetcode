package main

import "fmt"

func main() {
	left, right := 5, 7
	fmt.Printf("range bitwise AND: %d\n", rangeBitwiseAnd(left, right))
}

// https://leetcode.com/problems/bitwise-and-of-numbers-range
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}

	return right
}
