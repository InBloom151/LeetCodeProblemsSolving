### This approach has a time complexity of `O(n^3)`

### 1. Python

**Runtime:** `50 ms` faster than `99.62%` submissions  
**Memory usage:** `16.6 MB` less than `42.92%` submissions  

``` python
def four_sum(nums: list[int], target: int) -> list[list[int]]:
    nums.sort()
    results = []
    n = len(nums)
    for i in range(n - 3):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        if nums[i] * 4 > target:
            break
        if nums[i] + 3 * nums[-1] < target:
            continue
        for j in range(i + 1, n - 2):
            if j > i + 1 and nums[j] == nums[j - 1]:
                continue
            if nums[i] + nums[j] * 3 > target:
                break
            if nums[i] + nums[j] + 2 * nums[-1] < target:
                continue
            left, right = j + 1, n - 1
            while left < right:
                total = nums[i] + nums[j] + nums[left] + nums[right]
                if total == target:
                    results.append([nums[i], nums[j], nums[left], nums[right]])
                    left_val = nums[left]
                    while left < right and nums[left] == left_val:
                        left += 1
                    right_val = nums[right]
                    while left < right and nums[right] == right_val:
                        right -= 1
                elif total < target:
                    left += 1
                else:
                    right -= 1
    return results
```

### 2. TypeScript

**Runtime:** `75 ms` faster than `96.28%` submissions  
**Memory usage:** `56 MB` less than `40.56%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `4 ms` faster than `97.42%` submissions  
**Memory usage:** `2.8 MB` less than `80.88%` submissions  

``` go
import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	results := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]*4 > target {
			break
		}
		if nums[i]+3*nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]*3 > target {
				break
			}
			if nums[i]+nums[j]+2*nums[n-1] < target {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				total := nums[i] + nums[j] + nums[left] + nums[right]

				if total == target {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					leftVal := nums[left]
					for left < right && nums[left] == leftVal {
						left++
					}
					rightVal := nums[right]
					for left < right && nums[right] == rightVal {
						right--
					}
				} else if total < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return results
}
```