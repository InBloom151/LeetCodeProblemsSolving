def search_insert(nums: list[int], target: int) -> int:
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[left] < target:
            left += 1
        else:
            right -= 1
    return left

print(search_insert([1,3,5,6], 5))
print(search_insert([1,3,5,6], 2))
print(search_insert([1,3,5,6], 7))