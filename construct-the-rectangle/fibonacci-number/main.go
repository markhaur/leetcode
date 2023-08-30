package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("fibonacci number at %d is %d\n", n, fib(n))
}

// https://leetcode.com/problems/fibonacci-number/description
func fib(n int) int {
	if n == 0 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
