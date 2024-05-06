### This approach has a time complexity of `O(3^N * 4^M)`

### 1. Python

**Runtime:** `33 ms` faster than `72.44%` submissions  
**Memory usage:** `16.5 MB` less than `89.63%` submissions  

``` python
def letter_combinations(digits: str) -> list[str] | list:
    PHONE_NUMS = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
    if digits == "":
        return []
    
    result = list()
    get_combinations('', 0, digits, result, PHONE_NUMS)
    return result


def get_combinations(combination: str, index: int, digits: str, result: list, associations: dict) -> None:
    if index == len(digits):
        result.append(combination)
        return
    current_digit = digits[index]
    for letter in associations[current_digit]:
        get_combinations(combination + letter, index + 1, digits, result, associations)
```

### 2. TypeScript

**Runtime:** `52 ms` faster than `79.04%` submissions  
**Memory usage:** `51.3 MB` less than `63.88%` submissions  

``` typescript
function letterCombinations(digits: string): string[] {
    const PHONE_NUMS: { [key: string]: string } = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    };

    if (digits === "") {
        return [];
    }

    const result: string[] = [];
    getCombinations("", 0, digits, result, PHONE_NUMS);
    return result;
}

function getCombinations(
    combination: string,
    index: number,
    digits: string,
    result: string[],
    associations: { [key: string]: string }
): void {
    if (index === digits.length) {
        result.push(combination);
        return;
    }

    const currentDigit: string = digits[index];
    for (const letter of associations[currentDigit]) {
        getCombinations(combination + letter, index + 1, digits, result, associations);
    }
}
```

### 3. GO

**Runtime:** `1 ms` faster than `79.30%` submissions  
**Memory usage:** `2.3 MB` less than `26.60%` submissions  

``` go
func letterCombinations(digits string) []string {
	var PHONE_NUMS = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if digits == "" {
		return []string{}
	}

	result := []string{}
	getCombinations("", 0, digits, &result, PHONE_NUMS)
	return result

}

func getCombinations(combination string, index int, digits string, result *[]string, associations map[string]string) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	currentDigit := string(digits[index])
	letters := associations[currentDigit]
	for _, letter := range letters {
		getCombinations(combination+string(letter), index+1, digits, result, associations)
	}
}
```

### 4. Java

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `41.6 MB` less than `92.12%` submissions  

``` java
import java.util.ArrayList;
import java.util.List;

class solution {
    private static final String[] PHONE_NUMS = {
        "0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
    };
    
    public static List<String> letterCombinations(String digits) {
        List<String> result = new ArrayList<>();
        if (digits == null || digits.length() == 0) return result;
        
        backtrack(result, digits, 0, new StringBuilder());
        return result;
    }
    
    public static void backtrack(List<String> result, String digits, int index, StringBuilder combination) {
        if (index == digits.length()) {
            result.add(combination.toString());
            return;
        }
        
        int digit = digits.charAt(index) - '0';
        String letters = PHONE_NUMS[digit];
        for (char letter : letters.toCharArray()) {
            combination.append(letter);
            backtrack(result, digits, index + 1, combination);
            combination.deleteCharAt(combination.length() - 1);
        }
    }
}  
```