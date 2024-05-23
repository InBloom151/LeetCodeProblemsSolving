function searchRange(nums: number[], target: number): number[] {
    function findLeftMost(nums: number[], target: number): number {
        let [left, right] = [0, nums.length - 1];
        while (left <= right) {
            let mid = Math.floor((left + right) / 2);
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return left;
    }

    function findRightMost(nums: number[], target: number): number {
        let [left, right] = [0, nums.length - 1];
        while (left <= right) {
            let mid = Math.floor((left + right) / 2);
            if (nums[mid] <= target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return right;
    }

    let leftMost = findLeftMost(nums, target);
    let rightMost = findRightMost(nums, target);

    if (leftMost <= rightMost && leftMost < nums.length && nums[leftMost] === target) {
        return [leftMost, rightMost];
    } else {
        return [-1, -1];
    }
}

console.log(searchRange([5,7,7,8,8,10], 8));
console.log(searchRange([5,7,7,8,8,10], 6));
console.log(searchRange([], 0));