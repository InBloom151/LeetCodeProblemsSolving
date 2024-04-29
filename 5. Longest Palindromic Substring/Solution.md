### This approach has a time complexity of `O(n)`

Here we use a Manacher's algorithm.


### 1. Python

**Runtime:** `87 ms` faster than `96%` submissions  
**Memory usage:** `16.6 MB` less than `50.29%` submissions  

``` python
def preprocess(s: str) -> str:
    return '#' + '#'.join(s) + '#'

def palindrom_sub(s: str) -> str:
    if len(s) < 1:
        return ""

    processed_s = preprocess(s)
    n = len(processed_s)
    P = [0] * n
    C, R = 0, 0

    for i in range(1, n - 1):
        mirror = 2 * C - i
        
        if i < R:
            P[i] = min(R - i, P[mirror])
        
        while i + (1 + P[i]) < n and i - (1 + P[i]) >= 0 and processed_s[i + (1 + P[i])] == processed_s[i - (1 + P[i])]:
            P[i] += 1
        
        if i + P[i] > R:
            C = i
            R = i + P[i]
    
    max_len = max(P)
    center_index = P.index(max_len)
    
    start_index = (center_index - max_len) // 2
    return s[start_index:start_index + max_len]
```

### 2. TypeScript

**Runtime:** `94 ms` faster than `82.74%` submissions  
**Memory usage:** `56.1 MB` less than `69.60%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `3 ms` faster than `97.80%` submissions  
**Memory usage:** `4.8 MB` less than `62.37%` submissions  

``` go
import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}

	imin, imax, halfLen := 0, m, (m+n+1)/2
	var maxLeft, minRight int

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0
}
```

### 4. Java

**Runtime:** `1 ms` faster than `100.00%` submissions  
**Memory usage:** `45.5 MB` less than `91.31%` submissions  

``` java
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        int totalLength = m + n;
        
        if (totalLength % 2 == 1) {
            return findKthElement(nums1, nums2, totalLength / 2 + 1);
        } else {
            return (findKthElement(nums1, nums2, totalLength / 2) + findKthElement(nums1, nums2, totalLength / 2 + 1)) / 2.0;
        }
    }
    
    private int findKthElement(int[] nums1, int[] nums2, int k) {
        int m = nums1.length;
        int n = nums2.length;
        int index1 = 0, index2 = 0;
        
        while (true) {
            if (index1 == m) {
                return nums2[index2 + k - 1];
            }
            if (index2 == n) {
                return nums1[index1 + k - 1];
            }
            if (k == 1) {
                return Math.min(nums1[index1], nums2[index2]);
            }
            
            int halfK = k / 2;
            int newIndex1 = Math.min(index1 + halfK, m) - 1;
            int newIndex2 = Math.min(index2 + halfK, n) - 1;
            
            if (nums1[newIndex1] <= nums2[newIndex2]) {
                k -= (newIndex1 - index1 + 1);
                index1 = newIndex1 + 1;
            } else {
                k -= (newIndex2 - index2 + 1);
                index2 = newIndex2 + 1;
            }
        }
    }
}
```