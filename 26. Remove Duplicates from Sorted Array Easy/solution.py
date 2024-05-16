def remove_duplicates(nums: list[int]) -> int:
    count = 0
    for i in range(len(nums)):
        if i < len(nums) - 1 and nums[i] == nums[i + 1]:
            continue
        nums[count] = nums[i]
        count += 1
    return count


print(remove_duplicates([1,1,2]))
print(remove_duplicates([0,0,1,1,1,2,2,3,3,4]))