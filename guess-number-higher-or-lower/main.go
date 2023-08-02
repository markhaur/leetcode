package main

import "fmt"

var (
	GUESS = 0
)

func main() {
	GUESS = 7
	fmt.Printf("Guess number from 1 - 10: %d\n", guessNumber(10))
	GUESS = 8
	fmt.Printf("Guess number from 1 - 10: %d\n", guessNumber(10))
	GUESS = 2
	fmt.Printf("Guess number from 1 - 10: %d\n", guessNumber(10))
}

// https://leetcode.com/problems/guess-number-higher-or-lower/description
func guessNumber(n int) int {
	low, high := 0, n

	for low <= high {
		mid := low + (high-low)/2

		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

/**
 * the api method is provided by leetcode
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(g int) int {
	if g == GUESS {
		return 0
	}

	if g < GUESS {
		return 1
	}

	return -1
}
