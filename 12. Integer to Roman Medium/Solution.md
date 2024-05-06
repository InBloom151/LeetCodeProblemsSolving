### This approach has a time complexity of `O(1)`

### 1. Python

**Runtime:** `28 ms` faster than `99.70%` submissions  
**Memory usage:** `16.6 MB` less than `63.25%` submissions  

``` python
def int_to_roman(num: int) -> str:
    ASSOCIATIONS = {
        1000: "M",
        900: "CM",
        500: "D",
        400: "CD",
        100: "C",
        90: "XC",
        50: "L",
        40: "XL",
        10: "X",
        9: "IX",
        5: "V",
        4: "IV",
        1: "I",
    }

    result = ""

    for k, v in ASSOCIATIONS.items():
        while num >= k:
            result += v
            num -= k

    return result
```

### 2. TypeScript

**Runtime:** `102 ms` faster than `94.75%` submissions  
**Memory usage:** `57.3 MB` less than `40.66%` submissions  

``` typescript
function intToRoman(num: number): string {
    const romanNumerals: [number, string][] = [
        [1000, "M"],
        [900, "CM"],
        [500, "D"],
        [400, "CD"],
        [100, "C"],
        [90, "XC"],
        [50, "L"],
        [40, "XL"],
        [10, "X"],
        [9, "IX"],
        [5, "V"],
        [4, "IV"],
        [1, "I"],
    ];

    let result: string = '';

    for (let [value, roman] of romanNumerals) {
        while (num >= value) {
            result += roman;
            num -= value;
        }
        if (num === 0) break;
    }

    return result;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `5.3 MB` less than `15.22%` submissions  

``` go
func intToRoman(num int) string {

	var associations = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	result := strings.Builder{}

	for _, k := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= k {
			result.WriteString(associations[k])
			num -= k
		}
	}

	return result.String()
}
```

### 4. Java

**Runtime:** `3 ms` faster than `97.07%` submissions  
**Memory usage:** `43.9 MB` less than `94.41%` submissions  

``` java
class solution {
    public static String intToRoman(int num) {
        int[] nums = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] romans = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

        StringBuilder result = new StringBuilder();

        for (int i = 0; i < nums.length; i++) {
            int k = nums[i];
            String v = romans[i];
            while (num >= k) {
                result.append(v);
                num -= k;
            }
        }

        return result.toString();
    }
}
```