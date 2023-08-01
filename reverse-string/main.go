package main

import "fmt"

func main() {
	original := []byte{'h', 'e', 'l', 'l', 'o'}
	testArr := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(testArr)
	fmt.Printf("reverseString %v = %v\n", original, testArr)
}

// https://leetcode.com/problems/reverse-string/description
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for end-start > 0 {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
