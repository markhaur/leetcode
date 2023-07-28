package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("addDigits(%d) = %d\n", 38, addDigits(38))
	fmt.Printf("addDigits(%d) = %d\n", 0, addDigits(0))
}

// https://leetcode.com/problems/add-digits/description
func addDigits(num int) int {
	if num < 10 {
		return num
	}

	sum := 0
	size := int(math.Log10(float64(num)))
	divider := int(math.Pow10(size))

	for {
		if divider == 1 {
			sum += num
			return addDigits(sum)
		}

		d := num / divider
		sum += d
		num = num % divider
		divider = divider / 10
	}
}
