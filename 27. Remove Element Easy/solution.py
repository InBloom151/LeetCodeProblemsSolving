def remove_element(nums: list[int], val: int) -> int:
    n = len(nums) - 1
    for i in range(n, -1, -1):
        if nums[i] == val:
            nums.pop(i)
    return len(nums)

nums1 = [3,2,2,3]
print(remove_element(nums1, 3))
print(nums1)

nums2 = [0,1,2,2,3,0,4,2]
print(remove_element(nums2, 2))
print(nums2)