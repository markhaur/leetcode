package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Printf("Reverse vowels in %s => %s", s, reverseVowels(s))
}

// https://leetcode.com/problems/reverse-vowels-of-a-string/description
func reverseVowels(s string) string {
	start, end := 0, len(s)-1
	strBytes := []byte(s)
	for start < end {
		for start < end {
			if isVowel(strBytes[start]) {
				break
			}
			start++
		}

		for start < end {
			if isVowel(strBytes[end]) {
				strBytes[start], strBytes[end] = strBytes[end], strBytes[start]
				break
			}
			end--
		}

		start++
		end--
	}

	return string(strBytes)
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'A' ||
		c == 'e' || c == 'E' ||
		c == 'i' || c == 'I' ||
		c == 'o' || c == 'O' ||
		c == 'u' || c == 'U'
}
