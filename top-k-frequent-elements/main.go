package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8, 4, 5, 6, 7, 8}
	k := 8
	fmt.Printf("top k frequent: %v\n", topKFrequent(nums, k))
}

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int, 0)

	for _, num := range nums {
		store[num]++
	}

	return getMaxVals(k, store)
}

func getMaxVals(k int, store map[int]int) []int {
	maxVals := make([]int, k)
	topK := make([]int, k)

	for key, val := range store {
		for i := k - 1; i >= 0; i-- {
			if val > maxVals[i] {
				if i < k-1 {
					maxVals[i+1] = maxVals[i]
					topK[i+1] = topK[i]
				}
				maxVals[i] = val
				topK[i] = key
			} else {
				break
			}
		}
	}

	return topK
}
