package main

import (
	"fmt"
	"strings"
)

func main() {

	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Printf("word pattern: %t\n", wordPattern(pattern, s))
}

// https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	pMap := make(map[byte]string)
	sMap := make(map[string]byte)

	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		pVal, pFound := pMap[pattern[i]]
		sVal, sFound := sMap[strs[i]]

		if (!pFound && sFound) || (pFound && !sFound) {
			return false
		}

		if !pFound && !sFound {
			pMap[pattern[i]] = strs[i]
			sMap[strs[i]] = pattern[i]
			continue
		}

		if pVal != strs[i] || sVal != pattern[i] {
			return false
		}
	}
	return true
}
