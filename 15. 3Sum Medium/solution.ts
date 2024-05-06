function threeSum(nums: number[]): number[][] {
    nums.sort((a, b) => a - b);
    let n: number = nums.length;
    let result: any[] = [];

    for (let i = 0; i < n - 2; i++) {
        if (i > 0 && nums[i] == nums[i - 1]) continue;

        let left: number = i + 1;
        let right: number = n - 1;

        while (left < right) {
            let total: number = nums[i] + nums[left] + nums[right]

            if (total == 0) {
                result.push([nums[i], nums[left], nums[right]]);
                while (left < right && nums[left] == nums[left + 1]) left++;
                while (left < right && nums[right] == nums[right - 1]) right--;

                left++;
                right--;
            } else if (total < 0) {
                left++;
            } else {    
                right--;
            }
        }
    }

    return result;
}

console.log(threeSum([-1,0,1,2,-1,-4]))
console.log(threeSum([0,1,1]))
console.log(threeSum([0,0,0]))