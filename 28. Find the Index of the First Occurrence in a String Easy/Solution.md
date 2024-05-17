### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `29 ms` faster than `89.41%` submissions  
**Memory usage:** `16.5 MB` less than `42.86%` submissions  

``` python
def str_str(haystack: str, needle: str) -> int:
    return haystack.find(needle)
```

### 2. TypeScript

**Runtime:** `49 ms` faster than `91.32%` submissions  
**Memory usage:** `51.1 MB` less than `41.70%` submissions  

``` typescript
function strStr(haystack: string, needle: string): number {
    return haystack.indexOf(needle);
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.2 MB` less than `12.75%` submissions  

``` go
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

```