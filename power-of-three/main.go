package main

import "fmt"

func main() {
	fmt.Printf("isPowerOfThree(%d) = %t\n", 3, isPowerOfThree(3))
	fmt.Printf("isPowerOfThree(%d) = %t\n", 0, isPowerOfThree(0))
	fmt.Printf("isPowerOfThree(%d) = %t\n", -1, isPowerOfThree(-1))
	fmt.Printf("isPowerOfThree(%d) = %t\n", 1, isPowerOfThree(1))
	fmt.Printf("isPowerOfThree(%d) = %t\n", 81, isPowerOfThree(81))
}

func isPowerOfThree(n int) bool {
	power := 1
	for n >= power {
		if n == power {
			return true
		}
		power = power * 3
	}
	return false
}
