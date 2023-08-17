package main

import "fmt"

func main() {
	str := "Let's take LeetCode contest"
	fmt.Printf("Reverse words: %s", reverseWords(str))
}

// https://leetcode.com/problems/reverse-words-in-a-string-iii/description
func reverseWords(s string) string {
	arr := []byte(s)
	startIndex := 0
	endIndex := 0
	size := len(arr)

	for i := 0; i < size; i++ {
		if s[i] == ' ' || i+1 == size {
			endIndex = i - 1
			if i+1 == size {
				endIndex = i
			}

			for endIndex-startIndex > 0 {
				arr[startIndex], arr[endIndex] = arr[endIndex], arr[startIndex]
				startIndex++
				endIndex--
			}
			startIndex = i + 1
		}
	}

	return string(arr)
}
