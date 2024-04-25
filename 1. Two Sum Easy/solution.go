package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := make(map[int]int)

	for index, num := range nums {
		check := target - num
		if idx, found := result[check]; found {
			return []int{idx, index}
		}
		result[num] = index
	}

	return []int{}
}

func main() {
	examples := map[int][][]int{
		9: {{2, 7, 11, 15}, {0, 1}},
		6: {{3, 2, 4}, {1, 2}},
		8: {{4, 4}, {0, 1}},
	}

	for key, value := range examples {
		fmt.Println(equalSlices(twoSum(value[0], key), value[1]))
	}
}

func equalSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
