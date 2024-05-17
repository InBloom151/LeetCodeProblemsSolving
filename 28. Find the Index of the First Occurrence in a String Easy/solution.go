package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
