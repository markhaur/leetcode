package main

import "fmt"

func main() {
	fmt.Printf("isPowerOfFour(%d) = %t\n", 3, isPowerOfFour(4))
	fmt.Printf("isPowerOfFour(%d) = %t\n", 5, isPowerOfFour(5))
	fmt.Printf("isPowerOfFour(%d) = %t\n", -1, isPowerOfFour(-1))
	fmt.Printf("isPowerOfFour(%d) = %t\n", 1, isPowerOfFour(1))
	fmt.Printf("isPowerOfFour(%d) = %t\n", 0, isPowerOfFour(0))
}

// https://leetcode.com/problems/power-of-four/description
func isPowerOfFour(n int) bool {
	power := 1
	for n >= power {
		if n == power {
			return true
		}
		power = power * 4
	}
	return false
}
