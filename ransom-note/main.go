package main

import "fmt"

func main() {
	magazine := "baca"
	ransomNote := "aab"

	fmt.Printf("Can ransomnote `%s` be constructed from magazine `%s`: %t\n", ransomNote, magazine, canConstruct(ransomNote, magazine))
}

// https://leetcode.com/problems/ransom-note/description
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	charCount := make([]int, 26)

	for i, _ := range magazine {
		charCount[magazine[i]-'a']++
	}

	for i, _ := range ransomNote {
		if charCount[ransomNote[i]-'a'] == 0 {
			return false
		}
		charCount[ransomNote[i]-'a']--
	}

	return true
}
