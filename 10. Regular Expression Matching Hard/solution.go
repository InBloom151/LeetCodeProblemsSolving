package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
}

func isMatch(s string, p string) bool {
	memo := make(map[string]bool)

	var dp func(int, int) bool
	dp = func(i, j int) bool {
		key := fmt.Sprintf("%d,%d", i, j)
		if val, ok := memo[key]; ok {
			return val
		}

		var ans bool
		if j == len(p) {
			ans = i == len(s)
		} else {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2) || (firstMatch && dp(i+1, j))
			} else {
				ans = firstMatch && dp(i+1, j+1)
			}
		}

		memo[key] = ans
		return ans
	}

	return dp(0, 0)
}
