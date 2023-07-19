package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d is a happy number? %t\n", 19, isHappy(19))
	fmt.Printf("%d is a happy number? %t\n", 2, isHappy(2))
	fmt.Printf("%d is a happy number? %t\n", 18, isHappy(18))
	fmt.Printf("%d is a happy number? %t\n", 1, isHappy(1))
	fmt.Printf("%d is a happy number? %t\n", 125, isHappy(125))
	fmt.Printf("%d is a happy number? %t\n", 2147483647, isHappy(2147483647))
}

// https://leetcode.com/problems/happy-number/description/
func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	if n < 10 && n != 7 {
		return false
	}

	sum := 0
	size := int(math.Log10(float64(n)))
	divider := int(math.Pow10(size))

	for {
		if divider == 1 {
			sum += n * n
			return isHappy(sum)
		}

		d := n / divider
		sum += d * d
		n = n % divider
		divider = divider / 10
	}
}
