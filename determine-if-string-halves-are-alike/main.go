package main

import "fmt"

func main() {
	s := "book"
	fmt.Printf("is string halves are alike: %t\n", halvesAreAlike(s))
}

// https://leetcode.com/problems/determine-if-string-halves-are-alike
func halvesAreAlike(s string) bool {
	size := len(s)
	aCounter := 0
	bCounter := 0

	for i, z := 0, size/2; z < size; z++ {
		aCounter += isVowel(s[i])
		bCounter += isVowel(s[z])
		i++
	}

	return aCounter == bCounter
}

func isVowel(c byte) int {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return 1
	default:
		return 0
	}
}
