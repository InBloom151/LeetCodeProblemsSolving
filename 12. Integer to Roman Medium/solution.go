package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {

	var associations = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	result := strings.Builder{}

	for _, k := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= k {
			result.WriteString(associations[k])
			num -= k
		}
	}

	return result.String()
}
