### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `36 ms` faster than `89.16%` submissions  
**Memory usage:** `17.3 MB` less than `51.49` submissions  

``` python
def longest_valid_parentheses(s: str) -> int:
    stack = [-1]
    result = 0
    for i, c in enumerate(s):
        if c == "(":
            stack.append(i)
        else:
            stack.pop()
            if not stack:
                stack.append(i)
            else:
                result = max(result, i - stack[-1])
    return result
```

### 2. TypeScript

**Runtime:** `58 ms` faster than `84.13%` submissions  
**Memory usage:** `52.3 MB` less than `63.46%` submissions  

``` typescript
function longestValidParentheses(s: string): number {
    let stack: number[] = [-1];
    let result = 0;

    for (let i = 0; i < s.length; i++) {
        if (s.charAt(i) == "(") {
            stack.push(i);
        } else {
            stack.pop();
            if (stack.length == 0) {
                stack.push(i);
            } else {
                result = Math.max(result, i - stack[stack.length-1]);
            }
        }
    }

    return result;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `3.7 MB` less than `39.06%` submissions  

``` go
func longestValidParentheses(s string) int {
	stack := []int{-1}
	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = int(math.Max(float64(result), float64(i-stack[len(stack)-1])))
			}
		}
	}
	return result
}
```