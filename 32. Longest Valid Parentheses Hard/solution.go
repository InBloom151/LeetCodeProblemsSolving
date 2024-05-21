package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("((()))"))
	fmt.Println(longestValidParentheses("((()))()"))
	fmt.Println(longestValidParentheses("(()()())()((()))"))
}

func longestValidParentheses(s string) int {
	stack := []int{-1}
	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = int(math.Max(float64(result), float64(i-stack[len(stack)-1])))
			}
		}
	}
	return result
}
