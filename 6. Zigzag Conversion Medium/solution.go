package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, len(s))
	rowSize := 2*numRows - 2
	index := 0

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += rowSize {
			result[index] = s[j]
			index++
			if i != 0 && i != numRows-1 {
				diagonal := j + rowSize - 2*i
				if diagonal < len(s) {
					result[index] = s[diagonal]
					index++
				}
			}
		}
	}

	return string(result)
}
