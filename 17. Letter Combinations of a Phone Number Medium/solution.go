package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

func letterCombinations(digits string) []string {
	var PHONE_NUMS = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if digits == "" {
		return []string{}
	}

	result := []string{}
	getCombinations("", 0, digits, &result, PHONE_NUMS)
	return result

}

func getCombinations(combination string, index int, digits string, result *[]string, associations map[string]string) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	currentDigit := string(digits[index])
	letters := associations[currentDigit]
	for _, letter := range letters {
		getCombinations(combination+string(letter), index+1, digits, result, associations)
	}
}
