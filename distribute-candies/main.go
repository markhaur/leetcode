package main

import "fmt"

func main() {

	candyType := []int{1,1,2,2,3,3,9,3,4,6,7,8,11,12}
	fmt.Printf("distributed candies: %d\n", distributeCandies(candyType))
}

// https://leetcode.com/problems/distribute-candies/description
func distributeCandies(candyType []int) int {
	store := make(map[int]bool, 0)
	size := len(candyType)
	max := size / 2
	for _, ct := range candyType {
		store[ct] = true
		if len(store) >= max {
			return max
		}
	}

	return len(store)
}