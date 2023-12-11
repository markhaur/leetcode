package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	fmt.Printf("find special integer: %d\n", findSpecialInteger(arr))

}

// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array
func findSpecialInteger(arr []int) int {
	maxCount := -1
	maxElement := arr[0]
	counter := 1
	size := len(arr)
	for i := 1; i < size; i++ {
		if arr[i] == arr[i-1] {
			counter++
		} else {
			if counter > maxCount {
				maxCount = counter
				maxElement = arr[i-1]
			}
			counter = 1
		}
	}

	if counter > maxCount {
		return arr[size-1]
	}

	return maxElement
}
