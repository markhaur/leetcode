package main

import "fmt"

func main() {
	fmt.Printf("Pascal Triangle's row no %d => %v\n", 0, getRow(0))
	fmt.Printf("Pascal Triangle's row no %d => %v\n", 1, getRow(1))
	fmt.Printf("Pascal Triangle's row no %d => %v\n", 2, getRow(2))
	fmt.Printf("Pascal Triangle's row no %d => %v\n", 3, getRow(3))
	fmt.Printf("Pascal Triangle's row no %d => %v\n", 4, getRow(4))
}

/**
* Given an integer `rowIndex`, return the `rowIndexth` (0-indexed) row of the Pascal's triangle.
* In Pascal's triangle, each number is the sum of the two numbers directly above it.
 */
func getRow(rowIndex int) []int {
	row := []int{1}
	for i := 1; i <= rowIndex; i++ {
		var newRow = []int{1}
		for z := 1; z <= len(row); z++ {
			if i == z {
				newRow = append(newRow, 1)
			} else {
				newRow = append(newRow, row[z-1]+row[z])
			}
		}
		row = newRow
	}
	return row
}
