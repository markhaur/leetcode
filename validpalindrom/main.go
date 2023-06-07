package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputs := []string{"A man, a plan, a canal: Panama", "race a car", " "}

	for _, input := range inputs {
		fmt.Printf("%s is palindrome: %t\n", input, isPalindrome(input))
	}
}

func isPalindrome(s string) bool {
	s = removeNonAlphanumericChars(s)

	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	size := len(s)

	for i := 0; i < size/2; i++ {
		if s[i] != s[size-i-1] {
			return false
		}
	}

	return true
}

func removeNonAlphanumericChars(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	return reg.ReplaceAllString(s, "")
}
