from typing import List

def findMedianSortedArrays(nums1: List[int], nums2: List[int]) -> float:
    if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

    m, n = len(nums1), len(nums2)
    total_length = m + n
    half_length = total_length // 2

    left, right = 0, m
    while left < right:
        partition1 = left + (right - left) // 2
        partition2 = half_length - partition1

        if partition2 < 0:
            right = partition1 - 1
            continue
        if partition2 > n:
            left = partition1 + 1
            continue

        max_left1 = nums1[partition1 - 1] if partition1 > 0 else float("-inf")
        min_right1 = nums1[partition1] if partition1 < m else float("inf")
        max_left2 = nums2[partition2 - 1] if partition2 > 0 else float("-inf")
        min_right2 = nums2[partition2] if partition2 < n else float("inf")

        if max_left1 <= min_right2 and max_left2 <= min_right1:
            if total_length % 2 == 0:
                return (max(max_left1, max_left2) + min(min_right1, min_right2)) / 2
            else:
                return min(min_right1, min_right2)

        elif max_left1 > min_right2:
            right = partition1 - 1
        else:
            left = partition1 + 1

    partition1 = left
    partition2 = half_length - left
    max_left1 = nums1[partition1 - 1] if partition1 > 0 else float("-inf")
    min_right1 = nums1[partition1] if partition1 < m else float("inf")
    max_left2 = nums2[partition2 - 1] if partition2 > 0 else float("-inf")
    min_right2 = nums2[partition2] if partition2 < n else float("inf")

    if total_length % 2 == 0:
        return (max(max_left1, max_left2) + min(min_right1, min_right2)) / 2
    else:
        return min(min_right1, min_right2)


nums1 = [1, 2]
nums2 = [3, 4]
print(findMedianSortedArrays(nums1, nums2))
