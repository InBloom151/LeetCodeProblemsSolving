package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := math.Inf(1)
	minDiff := math.Inf(1)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			diff := math.Abs(float64(currentSum) - float64(target))

			if diff < minDiff {
				minDiff = diff
				closestSum = float64(currentSum)
			}

			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return target
			}
		}
	}

	return int(closestSum)
}
