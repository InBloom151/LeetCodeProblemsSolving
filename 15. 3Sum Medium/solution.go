package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[left] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return result
}
