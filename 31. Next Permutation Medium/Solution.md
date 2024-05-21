### This approach has a time complexity of `O(log(n))`

### 1. Python

**Runtime:** `39 ms` faster than `78.87%` submissions  
**Memory usage:** `16.5 MB` less than `93.13` submissions  

``` python
def next_permutation(nums: list[int]) -> None:
    i = len(nums) - 2
    while i >= 0 and nums[i] >= nums[i + 1]:
        i -= 1

    if i >= 0:
        j = len(nums) - 1
        while nums[j] <= nums[i]:
            j -= 1
        nums[i], nums[j] = nums[j], nums[i]
    
    left, right = i + 1, len(nums) - 1
    while left < right:
        nums[left], nums[right] = nums[right], nums[left]
        left += 1
        right -= 1
```

### 2. TypeScript

**Runtime:** `60 ms` faster than `73.96%` submissions  
**Memory usage:** `51.6 MB` less than `98.11%` submissions  

``` typescript
function nextPermutation(nums: number[]): void {
    let i = nums.length - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) i--;

    if (i >= 0) {
        let j = nums.length - 1;
        while (nums[j] <= nums[i]) j--;
        [nums[i], nums[j]] = [nums[j], nums[i]];
    }

    let [left, right] = [i + 1, nums.length - 1];

    while (left < right) {
        [nums[left], nums[right]] = [nums[right], nums[left]];
        left++;
        right--;
    }
}
```

### 3. GO

**Runtime:** `3 ms` faster than `61.35%` submissions  
**Memory usage:** `2.5 MB` less than `84.86%` submissions  

``` go
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
```