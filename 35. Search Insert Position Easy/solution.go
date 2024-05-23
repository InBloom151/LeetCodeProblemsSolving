package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] == target {
			return mid
		}

		if nums[left] < target {
			left++
		} else {
			right--
		}
	}

	return left
}
