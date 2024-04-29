package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(palindromSub("babad"))
	fmt.Println(palindromSub("cbbd"))
}

func preprocess(s string) string {
	return "#" + strings.Join(strings.Split(s, ""), "#") + "#"
}

func palindromSub(s string) string {
	if len(s) < 1 {
		return ""
	}

	processedS := preprocess(s)
	n := len(processedS)
	p := make([]int, n)
	c, r := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*c - i

		if i < r {
			p[i] = int(math.Min(float64(r-i), float64(p[mirror])))
		}

		for i+(1+p[i]) < n && i-(1+p[i]) >= 0 && processedS[i+(1+p[i])] == processedS[i-(1+p[i])] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0

	for i, num := range p {
		if num > maxLen {
			maxLen = num
			centerIndex = i
		}
	}

	startIdx := (centerIndex - maxLen) / 2
	return s[startIdx : startIdx+maxLen]
}
