package main

import "fmt"

func main() {
	fmt.Printf("%d is power of two? %t\n", 1, isPowerOfTwo(1))
	fmt.Printf("%d is power of two? %t\n", 16, isPowerOfTwo(16))
	fmt.Printf("%d is power of two? %t\n", 100, isPowerOfTwo(100))
}

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	two := 2
	for two <= n {
		if two == n {
			return true
		}
		two = two * 2
	}
	return false
}
