package main

import "fmt"

func main() {
	X, Y := 58, 405
	fmt.Printf("Hamming Distance between %d and %d is %d\n", X, Y, hammingDistance(X, Y))
}

// https://leetcode.com/problems/hamming-distance/description
func hammingDistance(x int, y int) int {
	distance := 0

	for x != 0 || y != 0 {
		if x%2 != y%2 {
			distance++
		}
		x = x / 2
		y = y / 2
	}

	return distance
}
