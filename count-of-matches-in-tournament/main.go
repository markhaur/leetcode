package main

import (
	"math" 
	"fmt"
)

func main() {
	n := 7
	fmt.Printf("number of matches: %d\n", numberOfMatches(n))
}

func numberOfMatches(n int) int {
	matchCount := 0
	for n != 1 {
		matchCount += (n / 2)
		n = int(math.Ceil(float64(n) / 2.0))
	}
	return matchCount
}
