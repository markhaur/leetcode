package main

import "fmt"

func main() {
	fmt.Printf("isPerfectSquare(%d) = %t\n", 4, isPerfectSquare(4))
	fmt.Printf("isPerfectSquare(%d) = %t\n", 5, isPerfectSquare(5))
	fmt.Printf("isPerfectSquare(%d) = %t\n", -1, isPerfectSquare(-1))
	fmt.Printf("isPerfectSquare(%d) = %t\n", 1, isPerfectSquare(1))
	fmt.Printf("isPerfectSquare(%d) = %t\n", 0, isPerfectSquare(0))
}

// https://leetcode.com/problems/valid-perfect-square/description
func isPerfectSquare(num int) bool {
	cmp := 0
	for i := 1; true; i++ {
		cmp = i * i
		if num == cmp {
			return true
		}
		if num < cmp {
			return false
		}
	}
	return false
}
