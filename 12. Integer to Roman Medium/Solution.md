### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `510 ms` faster than `80.31%` submissions  
**Memory usage:** `29.7 MB` less than `41.85%` submissions  

``` python
def max_area(height: list[int]) -> int:
    max_area = 0
    left = 0
    right = len(height) - 1

    while left < right:
        current_area = min(height[left], height[right]) * (right - left)
        max_area = max(max_area, current_area)

        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return max_area
```

### 2. TypeScript

**Runtime:** `66 ms` faster than `83.11%` submissions  
**Memory usage:** `58.4 MB` less than `99.15%` submissions  

``` typescript
function maxArea(height: number[]): number {
    let maxAreaValue = 0;
    let left = 0;
    let right = height.length - 1;

    while(left < right) {
        const currentArea = Math.min(height[left], height[right]) * (right - left);
        maxAreaValue = Math.max(maxAreaValue, currentArea);

        height[left] < height[right] ? ++left : --right;
    }

    return maxAreaValue;
}
```

### 3. GO

**Runtime:** `53 ms` faster than `93.63%` submissions  
**Memory usage:** `7.3 MB` less than `97.65%` submissions  

``` go
func maxArea(height []int) int {
    maxArea := 0
    left, right := 0, len(height)-1
    
    for left < right {
        area := min(height[left], height[right]) * (right - left)
        
        if area > maxArea {
            maxArea = area
        }
        
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### 4. Java

**Runtime:** `5 ms` faster than `71.01%` submissions  
**Memory usage:** `57.7 MB` less than `62.10%` submissions  

``` java
public class solution {
    public static int maxArea(int[] height) {
        int maxAreaValue = 0;
        int left = 0;
        int right = height.length - 1;

        while (left < right) {
            int currentArea = Math.min(height[left], height[right]) * (right - left);
            maxAreaValue = Math.max(maxAreaValue, currentArea);

            if (height[left] < height[right])
                left++;
            else
                right--;
        }

        return maxAreaValue;
    }
}
```