function nextPermutation(nums: number[]): void {
    let i = nums.length - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) i--;

    if (i >= 0) {
        let j = nums.length - 1;
        while (nums[j] <= nums[i]) j--;
        [nums[i], nums[j]] = [nums[j], nums[i]];
    }

    let [left, right] = [i + 1, nums.length - 1];

    while (left < right) {
        [nums[left], nums[right]] = [nums[right], nums[left]];
        left++;
        right--;
    }
}

let nums1 = [1,2,3];
nextPermutation(nums1);
console.log(nums1);

let nums2 = [3,2,1];
nextPermutation(nums2);
console.log(nums2);

let nums3 = [1,1,5];
nextPermutation(nums3);
console.log(nums3);