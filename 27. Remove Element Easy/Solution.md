### This approach has a time complexity of `O(n^2)`

### 1. Python

**Runtime:** `29 ms` faster than `94.20%` submissions  
**Memory usage:** `16.4 MB` less than `93.08%` submissions  

``` python
def remove_element(nums: list[int], val: int) -> int:
    n = len(nums) - 1
    for i in range(n, -1, -1):
        if nums[i] == val:
            nums.pop(i)
    return len(nums)
```

### 2. TypeScript

**Runtime:** `39 ms` faster than `99.06%` submissions  
**Memory usage:** `51.5 MB` less than `44.54%` submissions  

``` typescript
function removeElement(nums: number[], val: number): number {
    let i = 0;
    let j = 0;
    const n = nums.length;
    
    while (i < n) {
        if (nums[i] !== val) {
            nums[j] = nums[i];
            j++;
        }
        i++;
    }
    
    return j;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.4 MB` less than `67.96%` submissions  

``` go
func removeElement(nums []int, val int) int {
    j := 0
    for _, num := range nums {
        if num != val {
            nums[j] = num
            j++
        }
    }
    return j
}
```