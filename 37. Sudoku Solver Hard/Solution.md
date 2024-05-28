### This approach has a time complexity of `O(9^(n^2))`

### 1. Python

**Runtime:** `278 ms` faster than `41.98%` submissions  
**Memory usage:** `16.4 MB` less than `99.63%` submissions  

``` python
def solve_sudoku(board: list[list[str]]) -> None:
    def is_valid(row, col, num):
        for i in range(9):
            if board[i][col] == num or board[row][i] == num or board[3 * (row // 3) + i // 3][3 * (col // 3) + i % 3] == num:
                return False
        return True

    def backtrack(row=0, col=0):
        if row == 9:
            return True

        next_row = row + 1 if col == 8 else row
        next_col = (col + 1) % 9

        if board[row][col] != '.':
            return backtrack(next_row, next_col)

        for num in map(str, range(1, 10)):
            if is_valid(row, col, num):
                board[row][col] = num
                if backtrack(next_row, next_col):
                    return True
                board[row][col] = '.'

        return False

    backtrack()
```

### 2. TypeScript

**Runtime:** `69 ms` faster than `97.35%` submissions  
**Memory usage:** `52.9 MB` less than `71.97%` submissions  

``` typescript
function solveSudoku(board: string[][]): void {
    function isValid(row: number, col: number, num: string): boolean {
        for (let i = 0; i < 9; i++) {
            if (board[i][col] === num || board[row][i] === num || board[3 * Math.floor(row / 3) + Math.floor(i / 3)][3 * Math.floor(col / 3) + (i % 3)] === num) {
                return false;
            }
        }
        return true;
    }

    function backtrack(row: number = 0, col: number = 0): boolean {
        if (row === 9) {
            return true;
        }

        const nextRow = col === 8 ? row + 1 : row;
        const nextCol = (col + 1) % 9;

        if (board[row][col] !== '.') {
            return backtrack(nextRow, nextCol);
        }

        for (let num = 1; num <= 9; num++) {
            const strNum = num.toString();
            if (isValid(row, col, strNum)) {
                board[row][col] = strNum;
                if (backtrack(nextRow, nextCol)) {
                    return true;
                }
                board[row][col] = '.';
            }
        }

        return false;
    }

    backtrack();
}
```

### 3. GO

**Runtime:** `7 ms` faster than `32.07%` submissions  
**Memory usage:** `2.2 MB` less than `99.46%` submissions  

``` go
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
```