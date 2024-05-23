### This approach has a time complexity of `O(log n)`

### 1. Python

**Runtime:** `43 ms` faster than `90.61%` submissions  
**Memory usage:** `17.2 MB` less than `98.85%` submissions  

``` python
def search_insert(nums: list[int], target: int) -> int:
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[left] < target:
            left += 1
        else:
            right -= 1
    return left
```

### 2. TypeScript

**Runtime:** `46 ms` faster than `95.86%` submissions  
**Memory usage:** `51.8 MB` less than `8.94%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `3.1 MB` less than `72.57%` submissions  

``` go
import (
	"math"
)

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] == target {
			return mid
		}

		if nums[left] < target {
			left++
		} else {
			right--
		}
	}

	return left
}
```