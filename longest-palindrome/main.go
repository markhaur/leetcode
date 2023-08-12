package main

import "fmt"

func main() {
	s := "abcccckfdihfnxkcnhdfdfbjdndmfnvjbnASDFDFDFDFDDFddddd"
	fmt.Printf("longest palindrome: %d", longestPalindrome(s))
}

// https://leetcode.com/problems/longest-palindrome/description
func longestPalindrome(s string) int {
	wordCounter := make([]int, 52)

	for _, c := range s {
		if c-'A' < 26 {
			wordCounter[c-'A']++
		} else {
			wordCounter[c-'a'+26]++
		}
	}

	biggestOdd := 0
	evenSum := 0
	for _, wc := range wordCounter {
		if wc%2 == 0 {
			evenSum += wc
		} else {
			evenSum += wc - 1
			if biggestOdd < wc {
				biggestOdd = wc
			}
		}
	}

	if biggestOdd > 0 {
		return evenSum + 1
	}
	return evenSum
}
