function searchInsert(nums: number[], target: number): number {
    let [left, right] = [0, nums.length];
    while (left <= right) {
        let mid = Math.floor((left + right) / 2);
        if (nums[mid] === target) return mid;

        if (nums[left] < target) {
            left++;
        } else {
            right--;
        }
    }

    return left;
}

console.log(searchInsert([1,3,5,6], 5))
console.log(searchInsert([1,3,5,6], 2))
console.log(searchInsert([1,3,5,6], 7))