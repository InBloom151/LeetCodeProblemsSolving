### This approach has a time complexity of `O(n^2)`

### 1. Python

**Runtime:** `577 ms` faster than `84.88%` submissions  
**Memory usage:** `20.8 MB` less than `27.34%` submissions  

``` python
def three_sum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    n = len(nums)
    result = []

    for i in range(n - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue

        left = i + 1
        right = n - 1

        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total == 0:
                result.append([nums[i], nums[left], nums[right]])
                while left < right and nums[left] == nums[left + 1]:
                    left += 1
                while left < right and nums[right] == nums[right - 1]:
                    right -= 1
                left += 1
                right -= 1
            elif total < 0:
                left += 1
            else:
                right -= 1

    return result
```

### 2. TypeScript

**Runtime:** `140 ms` faster than `97.07%` submissions  
**Memory usage:** `64.8 MB` less than `97.26%` submissions  

``` typescript
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
```

### 3. GO

**Runtime:** `39 ms` faster than `85.36%` submissions  
**Memory usage:** `8.2 MB` less than `27.04%` submissions  

``` go
import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			total := nums[i] + nums[left] + nums[right]

			if total == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[left] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return result
}
```

### 4. Java

**Runtime:** `32 ms` faster than `57.11%` submissions  
**Memory usage:** `51 MB` less than `94.86%` submissions  

``` java
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;

class solution {
    public static List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        List<List<Integer>> result = new ArrayList<>();

        for (int i = 0; i < n - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;

            int left = i + 1;
            int right = n - 1;

            while(left < right) {
                int total = nums[i] + nums[left] + nums[right];

                if (total == 0) {
                    result.add(Arrays.asList(nums[i], nums[left], nums[right]));
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
}
```