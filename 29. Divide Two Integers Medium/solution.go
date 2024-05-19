package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}

func divide(dividend int, divisor int) int {
	MIN := -2147483648
	MAX := 2147483647

	if divisor == 0 || (dividend == MIN && divisor == -1) {
		return MAX
	}

	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend == MIN {
			return MAX
		}
		return -dividend
	}

	var sign int
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	} else {
		sign = 1
	}

	absoluteDividend := int(math.Abs(float64(dividend)))
	absoluteDivisor := int(math.Abs(float64(divisor)))

	quotient := 0

	for absoluteDividend >= absoluteDivisor {
		shift := 0
		tempDivisor := absoluteDivisor

		for absoluteDividend >= tempDivisor<<1 {
			shift++
			tempDivisor <<= 1
		}

		quotient += 1 << shift
		absoluteDividend -= tempDivisor
	}

	if sign == -1 {
		return -quotient
	}
	return quotient
}
