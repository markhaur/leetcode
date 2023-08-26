package main

import (
	"fmt"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Printf("findWords = %v\n", findWords(words))
}

// https://leetcode.com/problems/keyboard-row/description
func findWords(words []string) []string {
	keyboard := [][]byte{
		{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'},
		{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'},
		{'z', 'x', 'c', 'v', 'b', 'n', 'm'},
	}

	result := make([]string, 0)

	for _, word := range words {
		if isWordInRow(word, keyboard[0]) || isWordInRow(word, keyboard[1]) || isWordInRow(word, keyboard[2]) {
			result = append(result, word)
		}
	}

	return result
}

func isWordInRow(word string, row []byte) bool {
	found := false
	for i := 0; i < len(word); i++ {
		found = false
		for r := 0; r < len(row); r++ {
			if row[r]-word[i] == 0 || row[r]-word[i] == 32 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
