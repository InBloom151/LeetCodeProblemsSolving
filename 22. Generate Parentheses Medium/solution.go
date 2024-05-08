package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
	result := []string{}
	recGenerator("", 0, 0, n, &result)
	return result
}

func recGenerator(currentString string, opened int, closed int, n int, result *[]string) {
	if len(currentString) == 2*n {
		*result = append(*result, currentString)
		return
	}
	if opened < n {
		recGenerator(currentString+"(", opened+1, closed, n, result)
	}
	if closed < opened {
		recGenerator(currentString+")", opened, closed+1, n, result)
	}
}
