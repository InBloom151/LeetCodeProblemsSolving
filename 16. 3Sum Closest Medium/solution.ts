function threeSumClosest(nums: number[], target: number): number {
    nums.sort((a, b) => a - b);
    let closestSum = Infinity;
    let minDiff = Infinity;

    for (let i = 0; i < nums.length - 2; i++) {
        let left = i + 1;
        let right = nums.length - 1;

        while (left < right) {
            let currentSum = nums[i] + nums[left] + nums[right];
            let diff = Math.abs(currentSum - target);

            if (diff < minDiff) {
                minDiff = diff;
                closestSum = currentSum;
            }

            if (currentSum < target) {
                left++;
            } else if (currentSum > target) {
                right--;
            } else {
                return target;
            }
        }
    }

    return closestSum;
}

console.log(threeSumClosest([-1,2,1,-4], 1))
console.log(threeSumClosest([0,0,0], 1))