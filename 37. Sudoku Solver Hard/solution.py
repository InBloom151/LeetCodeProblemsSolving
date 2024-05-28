def solve_sudoku(board: list[list[str]]) -> None:
    def is_valid(row: int, col: int, num: str) -> bool:
        for i in range(9):
            if board[i][col] == num or board[row][i] == num or board[3 * (row // 3) + i // 3][3 * (col // 3) + i % 3] == num:
                return False
        return True

    def backtrack(row: int = 0, col: int = 0) -> bool:
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


board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]

solve_sudoku(board)

solved_board = [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

print(board == solved_board)