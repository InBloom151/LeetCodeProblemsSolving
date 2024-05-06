### This approach has a time complexity of `O(n^2)`

### 1. Python

**Runtime:** `344 ms` faster than `88.29%` submissions  
**Memory usage:** `16.6 MB` less than `59.59%` submissions  

``` python
def three_sum_closest(nums: list[int], target: int) -> int:
    nums.sort()
    closest_sum = float('inf')
    min_diff = float('inf')

    for i in range(len(nums) - 2):
        left = i + 1
        right = len(nums) - 1

        while left < right:
            current_sum = nums[i] + nums[left] + nums[right]
            diff = abs(current_sum - target)

            if diff < min_diff:
                min_diff = diff
                closest_sum = current_sum

            if current_sum < target:
                left += 1
            elif current_sum > target:
                right -= 1
            else:
                return target

    return closest_sum
```

### 2. TypeScript

**Runtime:** `70 ms` faster than `78.99%` submissions  
**Memory usage:** `52.3 MB` less than `33.73%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `11 ms` faster than `67.95%` submissions  
**Memory usage:** `3 MB` less than `70.88%` submissions  

``` go
import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := math.Inf(1)
	minDiff := math.Inf(1)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			diff := math.Abs(float64(currentSum) - float64(target))

			if diff < minDiff {
				minDiff = diff
				closestSum = float64(currentSum)
			}

			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return target
			}
		}
	}

	return int(closestSum)
}
```

### 4. Java

**Runtime:** `12 ms` faster than `96.39%` submissions  
**Memory usage:** `42.9 MB` less than `75.92%` submissions  

``` java
import java.util.Arrays;

class solution {
    public static int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        double closestSum = Double.POSITIVE_INFINITY;
        double minDiff = Double.POSITIVE_INFINITY;

        for (int i = 0; i < nums.length - 2; i++) {
            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                double currentSum = nums[i] + nums[left] + nums[right];
                double diff = Math.abs(currentSum - target);

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

        return (int) closestSum;
    }
}
```