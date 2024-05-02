### This approach has a time complexity of `O(log(n))`


### 1. Python

**Runtime:** `43 ms` faster than `91.70%` submissions  
**Memory usage:** `16.5 MB` less than `96.82%` submissions  

``` python
def is_palindrome(x: int) -> bool:
    return str(x) == ''.join(reversed(str(x)))
```

### 2. TypeScript

**Runtime:** `131 ms` faster than `74.05%` submissions  
**Memory usage:** `59.4 MB` less than `32.04%` submissions  

``` typescript
function isPalindrome(x: number): boolean {
    return x.toString() === x.toString().split("").reverse().join("");
}
```

### 3. GO

**Runtime:** `3 ms` faster than `97.76%` submissions  
**Memory usage:** `4.3 MB` less than `95.97%` submissions  

``` go
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}
```

### 4. Java

**Runtime:** `4 ms` faster than `100.00%` submissions  
**Memory usage:** `43.8 MB` less than `71.97%` submissions  

``` java
public class solution {
    public static boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }
        
        int reversed = 0;
        while (x > reversed) {
            reversed = reversed * 10 + x % 10;
            x /= 10;
        }
        
        return x == reversed || x == reversed / 10;
    }
}
```