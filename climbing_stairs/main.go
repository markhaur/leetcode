package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("There are %d ways to climb stairs wiht %d steps\n", climbStairs(i), i)
	}
}

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
func climbStairs(n int) int {
	var mem []int
	mem = append(mem, 1)
	mem = append(mem, 1)
	for i := 2; i <= n; i++ {
		mem = append(mem, mem[i-1]+mem[i-2])
	}

	return mem[n]
}
