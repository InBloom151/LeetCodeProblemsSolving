function twoSum(nums: number[], target: number): number[] {
    const result: { [key: number]: number } = {};
    
    for (let index = 0; index < nums.length; index++) {
        const num = nums[index];
        const check = target - num;
        if (result.hasOwnProperty(check)) {
            return [result[check], index];
        }
        result[num] = index;
    }
    
    return [];
}

const examples: { [key: number]: number[][] } = {
    9: [[2, 7, 11, 15], [0, 1]],
    6: [[3, 2, 4], [1, 2]],
    8: [[4, 4], [0, 1]]
};

for (const [key, value] of Object.entries(examples)) {
    console.log(JSON.stringify(twoSum(value[0], +key)) === JSON.stringify(value[1]));
}