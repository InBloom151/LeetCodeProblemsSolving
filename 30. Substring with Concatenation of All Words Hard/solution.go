package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	wordLength := len(words[0])
	wordCount := len(words)
	wordFrequency := map[string]int{}

	for _, word := range words {
		wordFrequency[word]++
	}

	result := []int{}

	for i := 0; i < wordLength; i++ {
		left := i
		right := i
		currentCount := 0
		currentFrequency := map[string]int{}

		for right+wordLength <= len(s) {
			word := s[right : right+wordLength]
			right += wordLength

			if wordFrequency[word] != 0 {
				currentFrequency[word]++
				currentCount++

				for currentFrequency[word] > wordFrequency[word] {
					leftWord := s[left : left+wordLength]
					currentFrequency[leftWord]--
					left += wordLength
					currentCount--
				}

				if currentCount == wordCount {
					result = append(result, left)
				}
			} else {
				currentFrequency = map[string]int{}
				currentCount = 0
				left = right
			}
		}
	}

	return result
}
