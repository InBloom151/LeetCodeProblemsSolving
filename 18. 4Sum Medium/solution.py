def four_sum(nums: list[int], target: int) -> list[list[int]]:
    nums.sort()
    results = []
    n = len(nums)
    for i in range(n - 3):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        if nums[i] * 4 > target:
            break
        if nums[i] + 3 * nums[-1] < target:
            continue
        for j in range(i + 1, n - 2):
            if j > i + 1 and nums[j] == nums[j - 1]:
                continue
            if nums[i] + nums[j] * 3 > target:
                break
            if nums[i] + nums[j] + 2 * nums[-1] < target:
                continue
            left, right = j + 1, n - 1
            while left < right:
                total = nums[i] + nums[j] + nums[left] + nums[right]
                if total == target:
                    results.append([nums[i], nums[j], nums[left], nums[right]])
                    left_val = nums[left]
                    while left < right and nums[left] == left_val:
                        left += 1
                    right_val = nums[right]
                    while left < right and nums[right] == right_val:
                        right -= 1
                elif total < target:
                    left += 1
                else:
                    right -= 1
    return results


print(four_sum([1,0,-1,0,-2,2], 0))
print(four_sum([2,2,2,2,2], 8))