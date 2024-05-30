package main

import (
	"fmt"
)

func main() {
	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	fmt.Println(combinationSum(candidates1, target1))
	candidates2 := []int{2, 3, 5}
	target2 := 8
	fmt.Println(combinationSum(candidates2, target2))
	candidates3 := []int{2}
	target3 := 1
	fmt.Println(combinationSum(candidates3, target3))
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	var backtrack func(remain int, combo []int, start int)
	backtrack = func(remain int, combo []int, start int) {
		if remain == 0 {
			combination := make([]int, len(combo))
			copy(combination, combo)
			result = append(result, combination)
			return
		} else if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			combo = append(combo, candidates[i])
			backtrack(remain-candidates[i], combo, i)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(target, []int{}, 0)
	return result
}
