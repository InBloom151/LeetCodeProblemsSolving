function removeElement(nums: number[], val: number): number {
    let i = 0;
    let j = 0;
    const n = nums.length;
    
    while (i < n) {
        if (nums[i] !== val) {
            nums[j] = nums[i];
            j++;
        }
        i++;
    }
    
    return j;
}

let nums1 = [3,2,2,3]
console.log(removeElement(nums1, 3))
console.log(nums1)

let nums2 = [0,1,2,2,3,0,4,2]
console.log(removeElement(nums2, 2))
console.log(nums2)