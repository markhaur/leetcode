package main

import "fmt"

func main() {
	s := "abcd"
	t := "adcbe"
	fmt.Printf("The difference in s->%s, t->%s is : %c\n", s, t, findTheDifference(s, t))
}

// https://leetcode.com/problems/find-the-difference/description
func findTheDifference(s string, t string) byte {
	var diff byte
	var i int
	for i = 0; i < len(s); i++ {
		diff += (t[i] - s[i])
	}
	diff += t[i]
	return diff
}
