function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    const mergedArray = mergeSortedArrays(nums1, nums2);
    const mergedLength = mergedArray.length;
    if (mergedLength % 2 === 0) {
        const mid = mergedLength / 2;
        return (mergedArray[mid - 1] + mergedArray[mid]) / 2;
    } else {
        return mergedArray[Math.floor(mergedLength / 2)];
    }
}

function mergeSortedArrays(nums1: number[], nums2: number[]): number[] {
    const merged: number[] = [];
    let i = 0, j = 0;

    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] < nums2[j]) {
            merged.push(nums1[i]);
            i++;
        } else {
            merged.push(nums2[j]);
            j++;
        }
    }

    while (i < nums1.length) {
        merged.push(nums1[i]);
        i++;
    }

    while (j < nums2.length) {
        merged.push(nums2[j]);
        j++;
    }

    return merged;
}

console.log(findMedianSortedArrays([1, 3], [2]));
console.log(findMedianSortedArrays([1, 2], [3, 4]));
