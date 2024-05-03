### This approach has a time complexity of `O(m*n)`

### IMPORTANT NOTE: In this problem, over 80% of the code was taken from various sources on the internet. Unfortunately, I was not able to solve this problem in optimal time by my own efforts


### 1. Python

**Runtime:** `67 ms` faster than `90.91%` submissions  
**Memory usage:** `52.9 MB` less than `51.78%` submissions  

``` python
def is_match(s: str, p: str) -> bool:
    memo = {}
    
    def dp(i: int, j: int) -> bool:
        if (i, j) in memo:
            return memo[(i, j)]

        if j == len(p):
            ans = i == len(s)
        else:
            first_match = i < len(s) and p[j] in {s[i], '.'}

            if j + 1 < len(p) and p[j + 1] == '*':
                ans = dp(i, j + 2) or (first_match and dp(i + 1, j))
            else:
                ans = first_match and dp(i + 1, j + 1)

        memo[(i, j)] = ans
        return ans

    return dp(0, 0)
```

### 2. TypeScript

**Runtime:** `131 ms` faster than `74.05%` submissions  
**Memory usage:** `59.4 MB` less than `32.04%` submissions  

``` typescript
function isMatch(s: string, p: string): boolean {
    const memo: {[key: string]: boolean} = {};

    function dp(i: number, j: number): boolean {
        const key = `${i},${j}`;
        if (memo[key] !== undefined) {
            return memo[key];
        }

        if (j === p.length) {
            const ans = i === s.length;
            memo[key] = ans;
            return ans;
        } else {
            const firstMatch = i < s.length && (p[j] === s[i] || p[j] === '.');

            if (j + 1 < p.length && p[j + 1] === '*') {
                const ans = dp(i, j + 2) || (firstMatch && dp(i + 1, j));
                memo[key] = ans;
                return ans;
            } else {
                const ans = firstMatch && dp(i + 1, j + 1);
                memo[key] = ans;
                return ans;
            }
        }
    }

    return dp(0, 0);
}
```

### 3. GO

**Runtime:** `3 ms` faster than `65.34%` submissions  
**Memory usage:** `3.3 MB` less than `12.65%` submissions  

``` go
func isMatch(s string, p string) bool {
	memo := make(map[string]bool)

	var dp func(int, int) bool
	dp = func(i, j int) bool {
		key := fmt.Sprintf("%d,%d", i, j)
		if val, ok := memo[key]; ok {
			return val
		}

		var ans bool
		if j == len(p) {
			ans = i == len(s)
		} else {
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2) || (firstMatch && dp(i+1, j))
			} else {
				ans = firstMatch && dp(i+1, j+1)
			}
		}

		memo[key] = ans
		return ans
	}

	return dp(0, 0)
}
```

### 4. Java

**Runtime:** `1 ms` faster than `100.00%` submissions  
**Memory usage:** `41.8 MB` less than `70.18%` submissions  

``` java
class solution {
    public static boolean isMatch(String s, String p) {
        int m = s.length();
        int n = p.length();
        boolean[][] memo = new boolean[m + 1][n + 1];
        memo[m][n] = true;

        for (int i = m; i >= 0; i--) {
            for (int j = n - 1; j >= 0; j--) {
                boolean firstMatch = i < m && (p.charAt(j) == s.charAt(i) || p.charAt(j) == '.');
                if (j + 1 < n && p.charAt(j + 1) == '*') {
                    memo[i][j] = memo[i][j + 2] || (firstMatch && memo[i + 1][j]);
                } else {
                    memo[i][j] = firstMatch && memo[i + 1][j + 1];
                }
            }
        }

        return memo[0][0];
    }

    public static void main(String[] args) {
        System.out.println(isMatch("aa", "a"));
        System.out.println(isMatch("aa", "a*"));
        System.out.println(isMatch("ab", ".*"));
    }
}
```