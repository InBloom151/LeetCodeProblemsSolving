def three_sum_closest(nums: list[int], target: int) -> int:
    nums.sort()
    closest_sum = float('inf')
    min_diff = float('inf')

    for i in range(len(nums) - 2):
        left = i + 1
        right = len(nums) - 1

        while left < right:
            current_sum = nums[i] + nums[left] + nums[right]
            diff = abs(current_sum - target)

            if diff < min_diff:
                min_diff = diff
                closest_sum = current_sum

            if current_sum < target:
                left += 1
            elif current_sum > target:
                right -= 1
            else:
                return target

    return closest_sum


print(three_sum_closest([-1,2,1,-4], 1))
print(three_sum_closest([0,0,0], 1))