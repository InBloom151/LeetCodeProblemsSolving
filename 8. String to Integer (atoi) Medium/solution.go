package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(" -042"))
	fmt.Println(myAtoi("1337c0d3"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("words and 987"))
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var result int
	for _, char := range s {
		if char < '0' || char > '9' {
			break
		}
		digit := int(char - '0')
		result = result*10 + digit

		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if result*sign < math.MinInt32 {
			return math.MinInt32
		}
	}

	return result * sign
}
