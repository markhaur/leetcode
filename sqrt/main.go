package main

import "fmt"

func main() {
	fmt.Printf("square root of 4 is %d", mySqrt(4))
	fmt.Printf("square root of 6 is %d", mySqrt(6))
	fmt.Printf("square root of 8 is %d", mySqrt(8))
	fmt.Printf("square root of 16 is %d", mySqrt(16))
}

// Given a non-negative integer x, return the square root of x
// rounded down to the nearest integer. The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
func mySqrt(x int) int {
	bucket := 0
	for num := 0; true; num += 2 {
		bucket = num * num
		if bucket > x {
			bucket = (num - 1) * (num - 1)
			if bucket > x {
				return num - 2
			}
			return num - 1
		}
	}
	return -1
}
