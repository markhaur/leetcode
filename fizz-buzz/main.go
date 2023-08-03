package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("fizzBuzz(15) => %v\n", fizzBuzz(15))
}

// https://leetcode.com/problems/fizz-buzz/description
func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch i % 15 {
		case 0:
			result[i-1] = "FizzBuzz"
		case 3, 6, 9, 12:
			result[i-1] = "Fizz"
		case 5, 10:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
