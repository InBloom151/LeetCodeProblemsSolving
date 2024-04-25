package main

func longestSubstringLength(s string) int {
	lastOccurrence := make([]int, 256)
	for i := range lastOccurrence {
		lastOccurrence[i] = -1
	}
	start := 0
	maxLength := 0

	for i, char := range s {
		charCode := int(char)
		if lastOccurrence[charCode] >= start {
			start = lastOccurrence[charCode] + 1
		}
		lastOccurrence[charCode] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	println(longestSubstringLength("abcabcbb"))
	println(longestSubstringLength("bbbbb"))
	println(longestSubstringLength("pwwkew"))
}
