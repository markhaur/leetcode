package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("There are %d ways to climb stairs wiht %d steps\n", climbStairs(i), i)
	}
}

func climbStairs(n int) int {
	var mem []int
	mem = append(mem, 1)
	mem = append(mem, 1)
	for i := 2; i <= n; i++ {
		mem = append(mem, mem[i-1]+mem[i-2])
	}

	return mem[n]
}
