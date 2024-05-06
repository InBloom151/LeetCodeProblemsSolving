package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			return first[:i]
		}
	}

	return first
}

Runtime: 2 ms, faster than 72.01% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.5 MB, less than 61.23% of Go online submissions for Longest Common Prefix.
