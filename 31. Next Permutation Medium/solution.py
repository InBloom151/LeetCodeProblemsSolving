def next_permutation(nums: list[int]) -> None:
    i = len(nums) - 2
    while i >= 0 and nums[i] >= nums[i + 1]:
        i -= 1

    if i >= 0:
        j = len(nums) - 1
        while nums[j] <= nums[i]:
            j -= 1
        nums[i], nums[j] = nums[j], nums[i]
    
    left, right = i + 1, len(nums) - 1
    while left < right:
        nums[left], nums[right] = nums[right], nums[left]
        left += 1
        right -= 1

nums1 = [1,2,3]
next_permutation(nums1)
print(nums1)

nums2 = [3,2,1]
next_permutation(nums2)
print(nums2)

nums3 = [1,1,5]
next_permutation(nums3)
print(nums3)