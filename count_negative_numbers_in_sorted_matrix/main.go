package main

import "fmt"

func main() {

	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{-1, -2, -3, -4},
	}

	fmt.Printf("Total negative integers in grid: %d\n", countNegatives(grid))

}

func countNegatives(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}
	matrixSize := len(grid[0])
	counter := 0
	for _, matrix := range grid {
		for i, val := range matrix {
			if val < 0 {
				counter += (matrixSize - i)
				break
			}
		}
	}
	return counter
}
