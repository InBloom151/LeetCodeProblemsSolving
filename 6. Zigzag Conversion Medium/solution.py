def convert(s: str, num_rows: int) -> str:
    if num_rows == 1 or num_rows >= len(s):
        return s
    
    result = [''] * num_rows
    row, step = 0, 1
    
    for char in s:
        result[row] += char
        if row == 0:
            step = 1
        elif row == num_rows - 1:
            step = -1
        row += step
    
    return ''.join(result)

print(convert("PAYPALISHIRING", 3))
print(convert("PAYPALISHIRING", 4))
print(convert("A", 1))