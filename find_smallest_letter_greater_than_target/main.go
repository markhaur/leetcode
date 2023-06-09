package main

import "fmt"

func main() {

	letters := []byte{'c', 'f', 'j'}
	target := 'a'
	fmt.Printf("next greatest letter: %d", nextGreatestLetter(letters, byte(target)))
}

// You are given an array of characters `letters` that is sorted in non-decreasing order, and a character `target`. There are at least two different characters in `letters`.
// Return the smallest character in `letters` that is lexicographically greater than `target`. If such a character does not exist, return the first character in `letters`.
func nextGreatestLetter(letters []byte, target byte) byte {

	if target > letters[len(letters)-1] {
		return letters[0]
	}

	for _, char := range letters {
		if char > target {
			return char
		}
	}

	return letters[0]
}
