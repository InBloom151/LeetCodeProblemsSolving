### This approach has a time complexity of `O(n)`


- Create an empty `hashmap` to store the indices of elements.
- Iterate through the array nums using a loop.
- For each element `num` in `nums`, calculate its complement, which is `target - num`.
- Check if the complement exists in the hashmap. If it does, it means we have found the two numbers that add up to the target. Return the indices of the current element and its complement from the hashmap.
- If the complement does not exist in the hashmap, add the current element and its index to the hashmap.
- If no two numbers are found that add up to the target, return an empty list.


### 1. Python

**Runtime:** `30 ms`  
**Memory usage:** `12.5 MB`

``` python
from typing import List

def twoSum(nums: List[int], target: int) -> List[int]:
    result = dict()
    
    for index, num in enumerate(nums):
        check = target - num
        if check in result:
            return [result[check], index]
        result[num] = index
        
    return []
```

### 2. TypeScript

**Runtime:** `57 ms`  
**Memory usage:** `52.3 MB`

``` typescript
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
```

### 3. GO

**Runtime:** `6 ms`  
**Memory usage:** `4.2 MB`

``` go
func twoSum(nums []int, target int) []int {
	result := make(map[int]int)

	for index, num := range nums {
		check := target - num
		if idx, found := result[check]; found {
			return []int{idx, index}
		}
		result[num] = index
	}

	return []int{}
}
```

### 4. Java

**Runtime:** `3 ms`  
**Memory usage:** `44.7 MB`

``` java
import java.util.HashMap;
import java.util.Map;

public class Main {

    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> result = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int check = target - nums[i];
            if (result.containsKey(check)) {
                return new int[]{result.get(check), i};
            }
            result.put(nums[i], i);
        }

        return new int[]{};
    }
}
```
