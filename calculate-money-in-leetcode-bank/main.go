package main

import "fmt"

func main() {
	n := 4
	fmt.Printf("total money: %d", totalMoney(n))
}

func totalMoney(n int) int {
	money := 0
	inc := 1
	for i := 0; i < n; i++ {
		money += (i % 7) + inc
		if (i+1)%7 == 0 {
			inc++
		}
	}

	return money
}
