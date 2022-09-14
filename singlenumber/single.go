package main

import "fmt"

func main() {
	var nums = []int{4, 1, 2, 1, 2}
	singleNum := singleNumber(nums)
	fmt.Printf("singleNumber: %v\n", singleNum)
}

func singleNumber(nums []int) int {
	singleNum := nums[0]

	for _, num := range nums[1:] {
		singleNum = singleNum ^ num
	}

	return singleNum
}
