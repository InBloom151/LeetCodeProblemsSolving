package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	isValid := func(row, col int, num byte) bool {
		for i := 0; i < 9; i++ {
			if board[i][col] == num || board[row][i] == num || board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
				return false
			}
		}
		return true
	}

	var backtrack func(row, col int) bool
	backtrack = func(row, col int) bool {
		if row == 9 {
			return true
		}

		nextRow := row
		nextCol := col + 1
		if nextCol == 9 {
			nextRow++
			nextCol = 0
		}

		if board[row][col] != '.' {
			return backtrack(nextRow, nextCol)
		}

		for num := byte('1'); num <= '9'; num++ {
			if isValid(row, col, num) {
				board[row][col] = num
				if backtrack(nextRow, nextCol) {
					return true
				}
				board[row][col] = '.'
			}
		}

		return false
	}

	backtrack(0, 0)
}
