function fourSum(nums: number[], target: number): number[][] {
    nums.sort((a, b) => a - b);
    let results: number[][] = [];
    let n = nums.length;
    for (let i = 0; i < n - 3; i++) {
        if (i > 0 && nums[i] === nums[i - 1]) continue;
        if (nums[i] * 4 > target) break;
        if (nums[i] + 3 * nums[n - 1] < target) continue;
        
        for (let j = i + 1; j < n - 2; j++) {
            if (j > i + 1 && nums[j] === nums[j - 1]) continue;
            if (nums[i] + nums[j] * 3 > target) break;
            if (nums[i] + nums[j] + 2 * nums[n - 1] < target) continue;
            let [left, right] = [j + 1, n - 1];

            while (left < right) {
                let total = nums[i] + nums[j] + nums[left] + nums[right];
                if (total === target) {
                    results.push([nums[i], nums[j], nums[left], nums[right]]);
                    let leftVal = nums[left];
                    while (left < right && nums[left] === leftVal) left++;
                    let rightVal = nums[right];
                    while (left < right && nums[right] === rightVal) right--; 
                } else if (total < target) {
                    left++;
                } else {
                    right--;
                }
            }
        }
    }

    return results;
}

console.log(fourSum([1,0,-1,0,-2,2], 0))
console.log(fourSum([2,2,2,2,2], 8))