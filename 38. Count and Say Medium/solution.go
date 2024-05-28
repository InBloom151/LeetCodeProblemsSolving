package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(1))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var nextSequence func(sequence string) string
	nextSequence = func(sequence string) string {
		result := []byte{}
		count := 1
		char := sequence[0]

		for i := 1; i < len(sequence); i++ {
			if sequence[i] == char {
				count++
			} else {
				result = append(result, []byte(strconv.Itoa(count))...)
				result = append(result, char)
				char = sequence[i]
				count = 1
			}
		}
		result = append(result, []byte(strconv.Itoa(count))...)
		result = append(result, char)

		return string(result)
	}

	currentSequence := "1"
	for i := 1; i < n; i++ {
		currentSequence = nextSequence(currentSequence)
	}

	return currentSequence
}
