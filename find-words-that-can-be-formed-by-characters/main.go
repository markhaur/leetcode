package main

import "fmt"

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Printf("characters count: %d\n", countCharacters(words, chars))
}

func countCharacters(words []string, chars string) int {
	alph := make([]int, 26)

	for _, ch := range chars {
		alph[ch-'a']++
	}

	sum := 0
	for _, word := range words {
		isFound := true
		nonFoundIdx := len(word)
		for idx, ch := range word {
			if alph[ch-'a'] == 0 {
				isFound = false
				nonFoundIdx = idx
				break
			}
			alph[ch-'a']--
		}
		if isFound {
			sum += len(word)
		}
		for idx, ch := range word {
			if idx >= nonFoundIdx {
				break
			}
			alph[ch-'a']++
		}
	}
	return sum
}
