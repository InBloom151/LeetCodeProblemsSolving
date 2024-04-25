### This approach has a time complexity of `O(n)`


### 1. Python

**Runtime:** `47 ms` faster than `89.08%` submissions  
**Memory usage:** `16.6 MB` less than `60.39%` submissions  

``` python
def longest_substring_length(s: str) -> int:
    last_occurrence = [-1] * 256
    start = 0
    max_length = 0

    for i, char in enumerate(s):
        char_code = ord(char)
        if last_occurrence[char_code] >= start:
            start = last_occurrence[char_code] + 1
        last_occurrence[char_code] = i
        max_length = max(max_length, i - start + 1)

    return max_length
```

### 2. TypeScript

**Runtime:** `62 ms` faster than `99.62%` submissions  
**Memory usage:** `54.4 MB` less than `84.19%` submissions  

``` typescript
function longestSubstringLength(s: string): number {
    const lastOccurrence: number[] = Array(256).fill(-1);
    let start = 0;
    let maxLength = 0;

    for (let i = 0; i < s.length; i++) {
        const charCode = s.charCodeAt(i);
        if (lastOccurrence[charCode] >= start) {
            start = lastOccurrence[charCode] + 1;
        }
        lastOccurrence[charCode] = i;
        maxLength = Math.max(maxLength, i - start + 1);
    }

    return maxLength;
}
```

### 3. GO

**Runtime:** `3 ms` faster than `89.79%` submissions  
**Memory usage:** `2.8 MB` less than `87.90%` submissions  

``` go
func longestSubstringLength(s string) int {
	lastOccurrence := make([]int, 256)
	for i := range lastOccurrence {
		lastOccurrence[i] = -1
	}
	start := 0
	maxLength := 0

	for i, char := range s {
		charCode := int(char)
		if lastOccurrence[charCode] >= start {
			start = lastOccurrence[charCode] + 1
		}
		lastOccurrence[charCode] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
```

### 4. Java

**Runtime:** `3 ms` faster than `90.48%%` submissions  
**Memory usage:** `43.4 MB` less than `87.57%` submissions  

``` java
import java.util.Arrays;

public class solution {
    public static int longestSubstringLength(String s) {
        int[] lastOccurrence = new int[256];
        Arrays.fill(lastOccurrence, -1);
        int start = 0;
        int maxLength = 0;

        for (int i = 0; i < s.length(); i++) {
            char currentChar = s.charAt(i);
            int charCode = (int) currentChar;
            if (lastOccurrence[charCode] >= start) {
                start = lastOccurrence[charCode] + 1;
            }
            lastOccurrence[charCode] = i;
            maxLength = Math.max(maxLength, i - start + 1);
        }

        return maxLength;
    }
}
```