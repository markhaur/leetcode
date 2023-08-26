package main

import "fmt"

func main() {
	fmt.Printf("%d complement is %d\n", 5, findComplement(5))
}

// https://leetcode.com/problems/number-complement/description
func findComplement(num int) int {
	complement := 0
	p := 1

	for num != 0 {
		if num%2 == 0 {
			complement += p
		}
		num /= 2
		p *= 2
	}

	return complement
}
