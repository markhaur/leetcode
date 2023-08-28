package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 1}}
	fmt.Printf("island perimeter of grid is %d\n", islandPerimeter(grid))
}

// https://leetcode.com/problems/island-perimeter/description
func islandPerimeter(grid [][]int) int {
	perimeter := 0
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for z := 0; z < cols; z++ {
			if grid[i][z] == 1 {
				perimeter += 4
				if i > 0 && grid[i-1][z] == 1 {
					perimeter -= 2
				}
				if z > 0 && grid[i][z-1] == 1 {
					perimeter -= 2
				}
			}
		}
	}

	return perimeter
}
