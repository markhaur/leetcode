package main

import "fmt"

func main() {
	colors := "abaac"
	neededTime := []int{1, 2, 3, 4, 5}
	fmt.Printf("min cost for removing baloons: %d\n", minCost(colors, neededTime))
}

// https://leetcode.com/problems/minimum-time-to-make-rope-colorful
func minCost(colors string, neededTime []int) int {
	totalBaloons := len(colors)
	minTime := 0

	i := 0
	for i < totalBaloons-1 {
		z := i + 1
		for z < totalBaloons {
			if colors[i] == colors[z] {
				if neededTime[i] < neededTime[z] {
					minTime += neededTime[i]
					i, z = z, z+1
				} else {
					minTime += neededTime[z]
					z++
				}
				continue
			}
			break
		}
		i = z
	}

	return minTime
}
