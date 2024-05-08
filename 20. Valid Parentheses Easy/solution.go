package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	BRACKETS := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	result := []rune{}

	for _, bracket := range s {
		if bracketIsOpen(bracket) {
			result = append(result, bracket)
		} else {
			if len(result) == 0 || result[len(result)-1] != BRACKETS[bracket] {
				return false
			}
			result = result[:len(result)-1]
		}
	}
	return len(result) == 0
}

func bracketIsOpen(char rune) bool {
	return char == '(' || char == '{' || char == '['
}
