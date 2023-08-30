package main

import "fmt"

func main() {
	area := 10001
	fmt.Printf("length and width for area %d is %v\n", area, constructRectangle(area))
}

// https://leetcode.com/problems/construct-the-rectangle/description
func constructRectangle(area int) []int {
	length, width := area, 1

	for i := 2; i*i <= area; i++ {
		if area%i == 0 {
			length = i
			width = area / i
		}
	}

	if width > length {
		length, width = width, length
	}

	return []int{length, width}
}