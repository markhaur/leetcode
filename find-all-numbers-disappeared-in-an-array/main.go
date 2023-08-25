package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Printf("disappeared numbers: %v\n", findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	filter := make([]int, len(nums))

	for _, n := range nums {
		filter[n-1] = -1
	}

	j := 0
	for i, _ := range filter {
		if filter[i] == 0 {
			filter[j] = i + 1
			j++
		}
	}

	return filter[:j]
}
