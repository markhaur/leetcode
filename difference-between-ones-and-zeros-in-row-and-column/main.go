package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}
	fmt.Printf("onesMinusZeros: %v\n", onesMinusZeros(grid))
}

// https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column
func onesMinusZeros(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	rowOnes, rowZeros := getRowInfo(grid)
	colOnes, colZeros := getColInfo(grid)

	for i := 0; i < rows; i++ {
		for z := 0; z < cols; z++ {
			grid[i][z] = rowOnes[i] + colOnes[z] - rowZeros[i] - colZeros[z]
		}
	}

	return grid
}

func getRowInfo(grid [][]int) ([]int, []int) {
	rows := len(grid)
	cols := len(grid[0])
	rowOnes := make([]int, rows)
	rowZeros := make([]int, rows)

	for i := 0; i < rows; i++ {
		oneCounter := 0
		for z := 0; z < cols; z++ {
			if grid[i][z] == 1 {
				oneCounter++
			}
		}
		rowOnes[i] = oneCounter
		rowZeros[i] = cols - oneCounter
	}

	return rowOnes, rowZeros
}

func getColInfo(grid [][]int) ([]int, []int) {
	rows := len(grid)
	cols := len(grid[0])
	colOnes := make([]int, cols)
	colZeros := make([]int, cols)

	for i := 0; i < cols; i++ {
		oneCounter := 0
		for z := 0; z < rows; z++ {
			if grid[z][i] == 1 {
				oneCounter++
			}
		}
		colOnes[i] = oneCounter
		colZeros[i] = rows - oneCounter
	}

	return colOnes, colZeros
}
