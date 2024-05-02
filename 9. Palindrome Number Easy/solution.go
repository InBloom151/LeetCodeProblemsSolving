package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}
