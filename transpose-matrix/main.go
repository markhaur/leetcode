package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("transpose = %v\n", transpose(matrix))
}

// https://leetcode.com/problems/transpose-matrix/
func transpose(matrix [][]int) [][]int {
	result := make([][]int, 0)
	rowSize := len(matrix)
	colSize := len(matrix[0])
	for i := 0; i < colSize; i++ {
		temp := make([]int, rowSize)
		for j := 0; j < rowSize; j++ {
			temp[j] = matrix[j][i]
		}
		result = append(result, temp)
	}
	return result
}
