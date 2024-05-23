### This approach has a time complexity of `O(log n)`

### 1. Python

**Runtime:** `71 ms` faster than `81.85%` submissions  
**Memory usage:** `17.8 MB` less than `84.42%` submissions  

``` python
def search_range(nums: list[int], target: int) -> list[int]:
    def find_left_most(nums: list[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return left

    def find_right_most(nums: list[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] <= target:
                left = mid + 1
            else:
                right = mid - 1
        return right

    leftmost = find_left_most(nums, target)
    rightmost = find_right_most(nums, target)
    
    if leftmost <= rightmost and leftmost < len(nums) and nums[leftmost] == target:
        return [leftmost, rightmost]
    else:
        return [-1, -1]
```

### 2. TypeScript

**Runtime:** `44 ms` faster than `98.98%` submissions  
**Memory usage:** `51.9 MB` less than `58.09%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `4.6 MB` less than `69.89%` submissions  

``` go
import (
	"math"
)

func findLeftMost(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func findRightMost(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func searchRange(nums []int, target int) []int {
	leftMost := findLeftMost(nums, target)
	rightMost := findRightMost(nums, target)

	if leftMost <= rightMost && leftMost < len(nums) && nums[rightMost] == target {
		return []int{leftMost, rightMost}
	} else {
		return []int{-1, -1}
	}
}
```