package main

import "fmt"

func main() {
	triangle := generate(5)

	fmt.Println("Pascal triangle is printed below")

	for _, row := range triangle {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// Given an integer `numRows`, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
// 		1
//     1 1
//    1 2 1
//   1 3 3 1
//  1 4 6 4 1
func generate(numRows int) [][]int {
	triangle := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		oldArr := triangle[i-1]
		var newArr = []int{1}
		for z := 1; z <= len(oldArr); z++ {
			if i == z {
				newArr = append(newArr, 1)
			} else {
				newArr = append(newArr, oldArr[z-1]+oldArr[z])
			}
		}
		triangle = append(triangle, newArr)
	}
	return triangle
}
