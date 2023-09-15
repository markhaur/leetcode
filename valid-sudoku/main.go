package main

import (
	"fmt"
	"sync"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Printf("vali sudoku: %t", isValidSudoku(board))
}

// https://leetcode.com/problems/valid-sudoku/description
func isValidSudoku(board [][]byte) bool {
	isLinesValid := false
	isBoxesValid := false
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		isLinesValid = validLines(board)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		subBox := [][]int{{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6}}
		for _, box := range subBox {
			isBoxesValid = validSubBox(box[0], box[1], board)
			if !isBoxesValid {
				return
			}
		}
	}()

	wg.Wait()

	return isLinesValid && isBoxesValid
}

func validLines(board [][]byte) bool {
	rowStore := make(map[byte]bool, 0)
	colStore := make(map[byte]bool, 0)

	for tracker := 0; tracker < 9; tracker++ {
		for z := 0; z < 9; z++ {
			if board[tracker][z] != '.' {
				if board[tracker][z] < 49 || board[tracker][z] > 57 {
					return false
				}
				if _, ok := rowStore[board[tracker][z]]; ok {
					return false
				}
				rowStore[board[tracker][z]] = true
			}
			if board[z][tracker] != '.' {
				if board[z][tracker] < 49 || board[z][tracker] > 57 {
					return false
				}
				if _, ok := colStore[board[z][tracker]]; ok {
					return false
				}
				colStore[board[z][tracker]] = true
			}
		}

		for key := range rowStore {
			delete(rowStore, key)
		}
		for key := range colStore {
			delete(colStore, key)
		}
	}

	return true
}

func validSubBox(rowS, colS int, board [][]byte) bool {
	store := make(map[byte]bool, 0)

	for i := rowS; i < rowS+3; i++ {
		for z := colS; z < colS+3; z++ {
			if board[i][z] != '.' {
				if board[i][z] < 49 || board[i][z] > 57 {
					return false
				}
				if _, ok := store[board[i][z]]; ok {
					return false
				}
				store[board[i][z]] = true
			}
		}
	}

	return true
}
