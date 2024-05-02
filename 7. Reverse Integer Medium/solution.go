package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
	isNegative := x < 0
	x = int(math.Abs(float64(x)))
	reversedNumber := 0

	for x > 0 {
		digit := x % 10
		reversedNumber = reversedNumber*10 + digit
		x /= 10
	}

	if reversedNumber > math.MaxInt32 {
		return 0
	}

	if isNegative {
		return -reversedNumber
	}
	return reversedNumber
}
