package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	ASSOCIATIONS := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	prevValue := 0
	var value int

	for numeral := len(s) - 1; numeral >= 0; numeral-- {
		value = ASSOCIATIONS[string(s[numeral])]

		if value < prevValue {
			total -= value
		} else {
			total += value
		}

		prevValue = value
	}

	return total
}

Runtime: 0 ms, faster than 100.00% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 83.70% of Go online submissions for Roman to Integer.
