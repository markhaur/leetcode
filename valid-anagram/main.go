package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Printf("s=%s and t=%s are valid anagrams: %t\n", s, t, isAnagram(s, t))
}

// https://leetcode.com/problems/valid-anagram/description
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)

	for i, _ := range s {
		count[s[i]-'a'] += 1
		count[t[i]-'a'] -= 1
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
