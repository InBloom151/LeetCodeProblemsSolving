package main

import "fmt"

func main() {
	board1 := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board1))

	board2 := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(board2))
}

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	subBoxes := make([][]map[byte]bool, 3)

	for i := range rows {
		rows[i] = make(map[byte]bool)
	}
	for i := range cols {
		cols[i] = make(map[byte]bool)
	}
	for i := range subBoxes {
		subBoxes[i] = make([]map[byte]bool, 3)
		for j := range subBoxes[i] {
			subBoxes[i][j] = make(map[byte]bool)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				if rows[i][num] {
					return false
				}
				rows[i][num] = true

				if cols[j][num] {
					return false
				}
				cols[j][num] = true

				if subBoxes[i/3][j/3][num] {
					return false
				}
				subBoxes[i/3][j/3][num] = true
			}
		}
	}
	return true
}
