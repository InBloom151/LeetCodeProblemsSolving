### This approach has a time complexity of `O(n*log(n))`

### 1. Python

**Runtime:** `31 ms` faster than `91.79%` submissions  
**Memory usage:** `16.7 MB` less than `44.13%` submissions  

``` python
def longest_common_prefix(strs: list[str]) -> str:
    if not strs:
            return ""

    strs.sort()
    first = strs[0]
    last = strs[-1]

    for i, char in enumerate(first):
        if char != last[i]:
            return first[:i]
    return first
```

### 2. TypeScript

**Runtime:** `51 ms` faster than `92.72%` submissions  
**Memory usage:** `51.8 MB` less than `65.05%` submissions  

``` typescript
function longestCommonPrefix(strs: string[]): string {
    if (!strs) return "";

    strs.sort();

    let first = strs[0];
    let last = strs[strs.length - 1];

    for(let i = 0; i < first.length; i++) {
        if (first[i] != last[i]) return first.substring(0, i);
    }

    return first;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.5 MB` less than `85.61%` submissions  

``` go
package main

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			return first[:i]
		}
	}

	return first
}
```

### 4. Java

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `41.3 MB` less than `61.43%` submissions  

``` java
import java.util.Arrays;

class solution {
    public static String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        }

        Arrays.sort(strs);

        String first = strs[0];
        String last = strs[strs.length - 1];

        for (int i = 0; i < first.length(); i++) {
            if (first.charAt(i) != last.charAt(i)) {
                return first.substring(0, i);
            }
        }

        return first;
    }
}
```