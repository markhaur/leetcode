package main

import "fmt"

func main() {
	bank := []string{"011001", "000000", "010100", "001000"}
	fmt.Printf("total number of beams: %d", numberOfBeams(bank))
}

// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
func numberOfBeams(bank []string) int {
	prev, totalBeams := 0, 0

	for _, floor := range bank {
		count := countDevices(floor)
		if count > 0 {
			totalBeams += (prev * count)
			prev = count
		}
	}

	return totalBeams
}

func countDevices(floor string) int {
	count := 0
	for _, device := range floor {
		if device == '1' {
			count++
		}
	}
	return count
}
