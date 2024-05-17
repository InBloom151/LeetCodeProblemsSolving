package main

import "fmt"

func main() {
	nums1 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(&nums1, 3))
	fmt.Println(nums1)

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(&nums2, 2))
	fmt.Println(nums2)
}

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	n := len(nums)

	for i < n {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	return j
}
