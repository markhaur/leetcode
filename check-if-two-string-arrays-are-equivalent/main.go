package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Printf("arrayStringsAreEqual: %t\n", arrayStringsAreEqual(word1, word2))
}

// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str := strings.Join(word1, "")
	size := len(str)

	counter := 0
	for _, w := range word2 {
		for i := 0; i < len(w); i++ {
			if size == counter || str[counter] != w[i] {
				return false
			}
			counter++
		}
	}

	return counter == size
}
