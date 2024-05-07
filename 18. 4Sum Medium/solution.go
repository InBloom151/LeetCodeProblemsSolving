package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]*4 > target {
			break
		}
		if nums[i]+3*nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]*3 > target {
				break
			}
			if nums[i]+nums[j]+2*nums[n-1] < target {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				total := nums[i] + nums[j] + nums[left] + nums[right]

				if total == target {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					leftVal := nums[left]
					for left < right && nums[left] == leftVal {
						left++
					}
					rightVal := nums[right]
					for left < right && nums[right] == rightVal {
						right--
					}
				} else if total < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return results
}
