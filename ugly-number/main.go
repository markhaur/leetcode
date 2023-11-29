package main

import "fmt"

func main() {
	fmt.Printf("%d is ugly number: %t\n", 1, isUgly(1))
	fmt.Printf("%d is ugly number: %t\n", 0, isUgly(0))
	fmt.Printf("%d is ugly number: %t\n", 2, isUgly(2))
	fmt.Printf("%d is ugly number: %t\n", 14, isUgly(14))
	fmt.Printf("%d is ugly number: %t\n", 55, isUgly(55))
	fmt.Printf("%d is ugly number: %t\n", 8, isUgly(8))
}

// https://leetcode.com/problems/ugly-number/
func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%2 == 0 {
		return isUgly(n / 2)
	}

	if n%3 == 0 {
		return isUgly(n / 3)
	}

	if n%5 == 0 {
		return isUgly(n / 5)
	}

	return false
}
