### This approach has a time complexity of `O(1)`

### 1. Python

**Runtime:** `82 ms` faster than `94.72%` submissions  
**Memory usage:** `16.5 MB` less than `94.36%` submissions  

``` python
def is_valid_sudoku(board):
    rows = [set() for _ in range(9)]
    cols = [set() for _ in range(9)]
    sub_boxes = [[set() for _ in range(3)] for _ in range(3)]

    for i in range(9):
        for j in range(9):
            num = board[i][j]
            if num != '.':
                if num in rows[i]:
                    return False
                rows[i].add(num)
                
                if num in cols[j]:
                    return False
                cols[j].add(num)
                
                if num in sub_boxes[i // 3][j // 3]:
                    return False
                sub_boxes[i // 3][j // 3].add(num)
                
    return True
```

### 2. TypeScript

**Runtime:** `72 ms` faster than `63.44%` submissions  
**Memory usage:** `53.4 MB` less than `67.76%` submissions  

``` typescript
function isValidSudoku(board: string[][]): boolean {
    const rows: Set<string>[] = Array.from({ length: 9 }, () => new Set());
    const cols: Set<string>[] = Array.from({ length: 9 }, () => new Set());
    const subBoxes: Set<string>[][] = Array.from({ length: 3 }, () => Array.from({ length: 3 }, () => new Set()));

    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            const num = board[i][j];
            if (num !== '.') {
                if (rows[i].has(num)) {
                    return false;
                }
                rows[i].add(num);

                if (cols[j].has(num)) {
                    return false;
                }
                cols[j].add(num);

                if (subBoxes[Math.floor(i / 3)][Math.floor(j / 3)].has(num)) {
                    return false;
                }
                subBoxes[Math.floor(i / 3)][Math.floor(j / 3)].add(num);
            }
        }
    }
    return true;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `3.9 MB` less than `26.06%` submissions  

``` go
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
```