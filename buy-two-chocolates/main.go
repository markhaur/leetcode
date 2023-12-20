package main

import "fmt"

func main() {
	prices := []int{1, 2, 2}
	money := 3
	fmt.Printf("leftover money after buying chocs: %d\n", buyChoco(prices, money))
}

// https://leetcode.com/problems/buy-two-chocolates/
func buyChoco(prices []int, money int) int {
	mins := []int{101, 101}

	for _, price := range prices {
		if price < mins[1] {
			mins[1] = price
		}

		if price < mins[0] {
			mins[0], mins[1] = mins[1], mins[0]
		}
	}
	leftover := money - mins[0] - mins[1]
	if leftover < 0 {
		return money
	}
	return leftover
}
