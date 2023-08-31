package main

import "fmt"

type Key [26]byte

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("group anagrams: %v\n", groupAnagrams(strs))
}

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	store := make(map[Key][]string, 0)

	for _, str := range strs {
		key := genKey(str)
		store[key] = append(store[key], str)
	}

	groups := make([][]string, 0)
	for _, v := range store {
		groups = append(groups, v)
	}
	return groups
}

func genKey(str string) (key Key) {

	for i := range str {
		key[str[i]-'a']++
	}

	return
}
