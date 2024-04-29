### This approach has a time complexity of `O(n)`

Here we use a Manacher's algorithm.


### 1. Python

**Runtime:** `87 ms` faster than `96%` submissions  
**Memory usage:** `16.6 MB` less than `50.29%` submissions  

``` python
def preprocess(s: str) -> str:
    return '#' + '#'.join(s) + '#'

def palindrome_sub(s: str) -> str:
    if len(s) < 1:
        return ""

    processed_s = preprocess(s)
    n = len(processed_s)
    P = [0] * n
    C, R = 0, 0

    for i in range(1, n - 1):
        mirror = 2 * C - i
        
        if i < R:
            P[i] = min(R - i, P[mirror])
        
        while i + (1 + P[i]) < n and i - (1 + P[i]) >= 0 and processed_s[i + (1 + P[i])] == processed_s[i - (1 + P[i])]:
            P[i] += 1
        
        if i + P[i] > R:
            C = i
            R = i + P[i]
    
    max_len = max(P)
    center_index = P.index(max_len)
    
    start_index = (center_index - max_len) // 2
    return s[start_index:start_index + max_len]
```

### 2. TypeScript

**Runtime:** `65 ms` faster than `99.58%` submissions  
**Memory usage:** `54.1 MB` less than `62.54%` submissions  

``` typescript
function preprocess(s: string): string {
    return "#" + s.split("").join("#") + "#";
}

function palindromeSub(s: string): string {
    if (s.length < 1) {
        return "";
    };

    let processedS = preprocess(s);
    let n = processedS.length;
    let p = new Array(n).fill(0);
    let c = 0;
    let r = 0;

    for (let i = 1; i < n - 1; i++) {
        let mirror = 2 * c - i;

        if (i < r) {
            p[i] = Math.min(r - i, p[mirror]);
        };

        while (i + p[i] < n && i - (1 + p[i]) >= 0 && processedS[i + (1 + p[i])] == processedS[i - (1 + p[i])]) {
            p[i] += 1;
        };

        if (i + p[i] > r) {
            c = i;
            r = i + p[i];
        };
    };

    let maxLen = Math.max.apply(null, p);
    let centerIndex = p.indexOf(maxLen);

    let startIndex = Math.ceil((centerIndex - maxLen) / 2);
    return s.slice(startIndex, startIndex + maxLen);
};
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `5.8 MB` less than `25.95%` submissions  

``` go
import (
	"math"
	"strings"
)


func preprocess(s string) string {
	return "#" + strings.Join(strings.Split(s, ""), "#") + "#"
}

func palindromSub(s string) string {
	if len(s) < 1 {
		return ""
	}

	processedS := preprocess(s)
	n := len(processedS)
	p := make([]int, n)
	c, r := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*c - i

		if i < r {
			p[i] = int(math.Min(float64(r-i), float64(p[mirror])))
		}

		for i+(1+p[i]) < n && i-(1+p[i]) >= 0 && processedS[i+(1+p[i])] == processedS[i-(1+p[i])] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0

	for i, num := range p {
		if num > maxLen {
			maxLen = num
			centerIndex = i
		}
	}

	startIdx := (centerIndex - maxLen) / 2
	return s[startIdx : startIdx+maxLen]
}
```

### 4. Java

**Runtime:** `11 ms` faster than `95.89%` submissions  
**Memory usage:** `43.1 MB` less than `65.85%` submissions  

``` java
public class Solution {

    public static String preprocess(String s) {
        StringBuilder processed = new StringBuilder("#");
        for (int i = 0; i < s.length(); i++) {
            processed.append(s.charAt(i)).append("#");
        }
        return processed.toString();
    }

    public static String palindromeSub(String s) {
        if (s.length() < 1) {
            return "";
        }

        String processedS = preprocess(s);
        int n = processedS.length();
        int[] P = new int[n];
        int C = 0, R = 0;

        for (int i = 1; i < n - 1; i++) {
            int mirror = 2 * C - i;
            if (i < R) {
                P[i] = Math.min(R - i, P[mirror]);
            }

            while (i + (1 + P[i]) < n && i - (1 + P[i]) >= 0 && processedS.charAt(i + (1 + P[i])) == processedS.charAt(i - (1 + P[i]))) {
                P[i]++;
            }

            if (i + P[i] > R) {
                C = i;
                R = i + P[i];
            }
        }

        int maxLen = 0;
        int centerIndex = 0;
        for (int i = 0; i < n; i++) {
            if (P[i] > maxLen) {
                maxLen = P[i];
                centerIndex = i;
            }
        }

        int startIndex = (centerIndex - maxLen) / 2;
        return s.substring(startIndex, startIndex + maxLen);
    }
}
```