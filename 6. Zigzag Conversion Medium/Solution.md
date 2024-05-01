### This approach has a time complexity of `O(n)`


### 1. Python

**Runtime:** `41 ms` faster than `95.08%` submissions  
**Memory usage:** `16.7 MB` less than `39.27%` submissions  

``` python
def convert(s: str, num_rows: int) -> str:
    if num_rows == 1 or num_rows >= len(s):
        return s
    
    result = [''] * num_rows
    row, step = 0, 1
    
    for char in s:
        result[row] += char
        if row == 0:
            step = 1
        elif row == num_rows - 1:
            step = -1
        row += step
    
    return ''.join(result)
```

### 2. TypeScript

**Runtime:** `76 ms` faster than `97.24%` submissions  
**Memory usage:** `56.5 MB` less than `68.25%` submissions  

``` typescript
function convert(s: string, numRows: number): string {
    if (numRows === 1 || numRows >= s.length) {
        return s;
    }

    const result = [];
    const stepSize = 2 * numRows - 2;
    const len = s.length;

    for (let i = 0; i < numRows; i++) {
        let index = i;
        let step1 = 2 * (numRows - 1 - i);
        let step2 = stepSize - step1;

        while (index < len) {
            if (step1 !== 0) {
                result.push(s[index]);
                index += step1;
            }
            if (step2 !== 0 && index < len) {
                result.push(s[index]);
                index += step2;
            }
        }
    }

    return result.join("");
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `4 MB` less than `91.24%` submissions  

``` go
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, len(s))
	rowSize := 2*numRows - 2
	index := 0

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += rowSize {
			result[index] = s[j]
			index++
			if i != 0 && i != numRows-1 {
				diagonal := j + rowSize - 2*i
				if diagonal < len(s) {
					result[index] = s[diagonal]
					index++
				}
			}
		}
	}

	return string(result)
}
```

### 4. Java

**Runtime:** `2 ms` faster than `99.44%` submissions  
**Memory usage:** `44.7 MB` less than `66.69%` submissions  

``` java
class Solution {
    public static String convert(String s, int numRows) {
        if (numRows >= s.length() || numRows == 1) {
            return s;
        }

        char[] result = new char[s.length()];
        int rowSize = 2 * numRows - 2;
        int index = 0;

        for (int i = 0; i < numRows; i++) {
            for (int j = i; j < s.length(); j += rowSize) {
                result[index++] = s.charAt(j);
                if (i != 0 && i != numRows - 1) {
                    int diagonal = j + rowSize - 2 * i;
                    if (diagonal < s.length()) {
                        result[index++] = s.charAt(diagonal);
                    }
                }
            }
        }
        return new String(result);
    }
}
```