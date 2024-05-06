### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `39 ms` faster than `89.11%` submissions  
**Memory usage:** `16.6 MB` less than `69.79%` submissions  

``` python
def roman_to_int(s: str) -> int:
    ASSOCIATIONS = {
        'I': 1, 
        'V': 5, 
        'X': 10, 
        'L': 50,
        'C': 100, 
        'D': 500, 
        'M': 1000,
    }

    total = 0
    prev_value = 0

    for numeral in reversed(s):
        value = ASSOCIATIONS[numeral]
        if value < prev_value:
            total -= value
        else:
            total += value
        prev_value = value

    return total
```

### 2. TypeScript

**Runtime:** `104 ms` faster than `92.52%` submissions  
**Memory usage:** `55.7 MB` less than `64.26%` submissions  

``` typescript
function romanToInt(s: string): number {
    const ASSOCIATIONS: {[key: string]: number} = {
        'I': 1, 
        'V': 5, 
        'X': 10, 
        'L': 50,
        'C': 100, 
        'D': 500, 
        'M': 1000,
    }

    let total = 0;
    let prevValue = 0;

    for (let numeral = s.length - 1; numeral >= 0; numeral--) {
        let value = ASSOCIATIONS[s[numeral]];

        total = value < prevValue ? total -= value : total += value;

        prevValue = value;
    }

    return total;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `3 MB` less than `83.70%` submissions  

``` go
func romanToInt(s string) int {
	ASSOCIATIONS := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	prevValue := 0
	var value int

	for numeral := len(s) - 1; numeral >= 0; numeral-- {
		value = ASSOCIATIONS[string(s[numeral])]

		if value < prevValue {
			total -= value
		} else {
			total += value
		}

		prevValue = value
	}

	return total
}
```

### 4. Java

**Runtime:** `4 ms` faster than `68.28%` submissions  
**Memory usage:** `44.6 MB` less than `61.44%` submissions  

``` java
import java.util.Map;
import java.util.HashMap;

class solution {
    public static int romanToInt(String s) {
        Map<String, Integer> ASSOCIATIONS = new HashMap<>();
        ASSOCIATIONS.put("I", 1);
        ASSOCIATIONS.put("V", 5);
        ASSOCIATIONS.put("X", 10);
        ASSOCIATIONS.put("L", 50);
        ASSOCIATIONS.put("C", 100);
        ASSOCIATIONS.put("D", 500);
        ASSOCIATIONS.put("M", 1000);
    
        int total = 0;
        int prevValue = 0;
    
        for (int numeral = s.length() - 1; numeral >= 0; numeral--) {
            int value = ASSOCIATIONS.get(String.valueOf(s.charAt(numeral)));
    
            total = value < prevValue ? total - value : total + value;
    
            prevValue = value;
        }
    
        return total;
    }
}
```

Runtime: 4 ms, faster than 68.26% of Java online submissions for Roman to Integer.
Memory Usage: 44.6 MB, less than 61.44% of Java online submissions for Roman to Integer.