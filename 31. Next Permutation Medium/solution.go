package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nextPermutation(nums1)
	fmt.Println(nums1)

	nums2 := []int{3, 2, 1}
	nextPermutation(nums2)
	fmt.Println(nums2)

	nums3 := []int{1, 1, 5}
	nextPermutation(nums3)
	fmt.Println(nums3)
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
