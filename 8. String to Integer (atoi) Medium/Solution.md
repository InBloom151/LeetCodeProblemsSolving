### This approach has a time complexity of `O(n)`


### 1. Python

**Runtime:** `35 ms` faster than `74.45%` submissions  
**Memory usage:** `16.6 MB` less than `32.32%` submissions  

``` python
def myAtoi(s: str) -> int:
    s = s.strip()
    
    if not s:
        return 0
    
    sign = 1
    if s[0] == '-':
        sign = -1
        s = s[1:]
    elif s[0] == '+':
        s = s[1:]
    
    result = 0
    for char in s:
        if not char.isdigit():
            break
        result = result * 10 + int(char)
    
    result *= sign
    
    INT_MAX = 2**31 - 1
    INT_MIN = -2**31
    if result > INT_MAX:
        return INT_MAX
    elif result < INT_MIN:
        return INT_MIN
    
    return result
```

### 2. TypeScript

**Runtime:** `80 ms` faster than `61.23%` submissions  
**Memory usage:** `55.2 MB` less than `9.75%` submissions  

``` typescript
function myAtoi(s: string): number {
    let sign = 1, result = 0, i = 0;
    const INT_MAX = 2 ** 31 - 1, INT_MIN = -(2 ** 31);

    while (s[i] === ' ') i++;
    if (s[i] === '-') { sign = -1; i++; }
    else if (s[i] === '+') i++;

    for (; i < s.length && '0' <= s[i] && s[i] <= '9'; i++) {
        result = result * 10 + (s.charCodeAt(i) - 48);
        if (sign * result > INT_MAX) return INT_MAX;
        if (sign * result < INT_MIN) return INT_MIN;
    }

    return sign * result;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.4 MB` less than `79.72%` submissions  

``` go
package main

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var result int
	for _, char := range s {
		if char < '0' || char > '9' {
			break
		}
		digit := int(char - '0')
		result = result*10 + digit

		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if result*sign < math.MinInt32 {
			return math.MinInt32
		}
	}

	return result * sign
}
```

### 4. Java

**Runtime:** `1 ms` faster than `100.00%` submissions  
**Memory usage:** `42.3 MB` less than `60.09%` submissions  

``` java
public static int myAtoi(String s) {
    int index = 0, sign = 1, result = 0;
    while (index < s.length() && s.charAt(index) == ' ')
        index++;
    if (index < s.length() && (s.charAt(index) == '-' || s.charAt(index) == '+')) {
        sign = (s.charAt(index) == '-') ? -1 : 1;
        index++;
    }
    while (index < s.length() && Character.isDigit(s.charAt(index))) {
        int digit = s.charAt(index) - '0';
        if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && digit > 7)) {
            return (sign == 1) ? Integer.MAX_VALUE : Integer.MIN_VALUE;
        }
        result = result * 10 + digit;
        index++;
    }
    return result * sign;
}
```