package main

import "fmt"

func main() {
	s := "abc"
	t := "kdfhdkhakdjfbkdfc"
	fmt.Printf("is subsequence: %t\n", isSubsequence(s, t))
}

// https://leetcode.com/problems/is-subsequence/
func isSubsequence(s string, t string) bool {
	sSize := len(s)
	tSize := len(t)

	if tSize < sSize {
		return false
	}

	z := 0
	found := false
	for i := 0; i < sSize; i++ {
		found = false
		for z < tSize {
			if s[i] == t[z] {
				found = true
				z++
				break
			}
			z++
		}
		if !found {
			return false
		}
	}
	return true
}
