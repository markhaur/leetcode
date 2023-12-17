package main

import "fmt"

func main() {
	fmt.Printf("can win nim with 4 stones: %t\n", canWinNim(4))
}

// https://leetcode.com/problems/nim-game/description
func canWinNim(n int) bool {
	return n%4 != 0
}
