package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Printf("content children count: %d\n", findContentChildren(g, s))
}

// https://leetcode.com/problems/assign-cookies
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	contentChilds := 0
	cookie := 0
	greed := 0
	gSize := len(g)
	sSize := len(s)

	for greed < gSize && cookie < sSize {
		if s[cookie] >= g[greed] {
			contentChilds++
			greed++
		}
		cookie++
	}

	return contentChilds
}
