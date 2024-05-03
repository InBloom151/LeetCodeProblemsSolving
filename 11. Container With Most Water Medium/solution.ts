function maxArea(height: number[]): number {
    let maxAreaValue = 0;
    let left = 0;
    let right = height.length - 1;

    while(left < right) {
        const currentArea = Math.min(height[left], height[right]) * (right - left);
        maxAreaValue = Math.max(maxAreaValue, currentArea);

        height[left] < height[right] ? ++left : --right;
    }

    return maxAreaValue;
}

console.log(maxArea([1,8,6,2,5,4,8,3,7]));
console.log(maxArea([1,1]));