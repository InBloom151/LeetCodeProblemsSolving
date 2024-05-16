### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `73 ms` faster than `71.19%` submissions  
**Memory usage:** `18 MB` less than `57.75%` submissions  

``` python
def remove_duplicates(nums: list[int]) -> int:
    count = 0
    for i in range(len(nums)):
        if i < len(nums) - 1 and nums[i] == nums[i + 1]:
            continue
        nums[count] = nums[i]
        count += 1
    return count
```

### 2. TypeScript

**Runtime:** `57 ms` faster than `94.55%` submissions  
**Memory usage:** `52.7 MB` less than `73.70%` submissions  

``` typescript
function removeDuplicates(nums: number[]): number {
    let count = 0;
    for (let i = 0; i < nums.length; i++) {
        if (i < nums.length - 1 && nums[i] == nums[i + 1]) {continue}

        nums[count] = nums[i];
        count++;
    }
    return count;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `4.5 MB` less than `93.75%` submissions  

``` go
func removeDuplicates(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}
		nums[count] = nums[i]
		count++
	}
	return count
}
```