package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func findLeftMost(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func findRightMost(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func searchRange(nums []int, target int) []int {
	leftMost := findLeftMost(nums, target)
	rightMost := findRightMost(nums, target)

	if leftMost <= rightMost && leftMost < len(nums) && nums[rightMost] == target {
		return []int{leftMost, rightMost}
	} else {
		return []int{-1, -1}
	}
}
