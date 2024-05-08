### This approach has a time complexity of `O(n)`

### 1. Python

**Runtime:** `29 ms` faster than `90.50%` submissions  
**Memory usage:** `16.5 MB` less than `98.05%` submissions  

``` python
def is_valid(s: str) -> bool:
    BRACKETS = {
        ")": "(",
        "]": "[",
        "}": "{",
    }
    result = []

    for bracket in s:
        if bracket in BRACKETS:
            top_char = result.pop() if result else "#"
            if BRACKETS[bracket] != top_char:
                return False
        else:
            result.append(bracket)
    return not result
```

### 2. TypeScript

**Runtime:** `60 ms` faster than `69.36%` submissions  
**Memory usage:** `52.4 MB` less than `64.31%` submissions  

``` typescript
function isValid(s: string): boolean {
    const result: string[] = [];
    const openBrackets = ['(', '[', '{'];
    const closingBrackets = [')', ']', '}'];

    for (let bracket of s) {
        if (openBrackets.includes(bracket)) {
            result.push(bracket);
        } else if (closingBrackets.includes(bracket)) {
            const expectedBracket = openBrackets[closingBrackets.indexOf(bracket)];
            if (result.pop() !== expectedBracket) return false;
        }
    }

    return result.length === 0;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.2 MB` less than `63.36%` submissions  

``` go
func isValid(s string) bool {
	BRACKETS := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	result := []rune{}

	for _, bracket := range s {
		if bracketIsOpen(bracket) {
			result = append(result, bracket)
		} else {
			if len(result) == 0 || result[len(result)-1] != BRACKETS[bracket] {
				return false
			}
			result = result[:len(result)-1]
		}
	}
	return len(result) == 0
}

func bracketIsOpen(char rune) bool {
	return char == '(' || char == '{' || char == '['
}
```