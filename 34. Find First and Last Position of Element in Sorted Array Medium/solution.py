def search_range(nums: list[int], target: int) -> list[int]:
    def find_left_most(nums: list[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return left

    def find_right_most(nums: list[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] <= target:
                left = mid + 1
            else:
                right = mid - 1
        return right

    leftmost = find_left_most(nums, target)
    rightmost = find_right_most(nums, target)
    
    if leftmost <= rightmost and leftmost < len(nums) and nums[leftmost] == target:
        return [leftmost, rightmost]
    else:
        return [-1, -1]


        

print(search_range([5,7,7,8,8,10], 8))
print(search_range([5,7,7,8,8,10], 6))
print(search_range([], 0))