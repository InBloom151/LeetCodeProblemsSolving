### This approach has a time complexity of `O(2^n)`

### 1. Python

**Runtime:** `33 ms` faster than `83.05%` submissions  
**Memory usage:** `16.6 MB` less than `99.96%` submissions  

``` python
def generate_parenthesis(n: int) -> list[str]:
    result = []
    rec_generator("", 0, 0, n, result)
    return result

def rec_generator(current_string: str, opened: int, closed: int, n: int, result: list) -> None:
    if len(current_string) == 2 * n:
        result.append(current_string)
        return
    if opened < n:
        rec_generator(current_string + '(', opened + 1, closed, n, result)
    if closed < opened:
        rec_generator(current_string + ')', opened, closed + 1, n, result)
```

### 2. TypeScript

**Runtime:** `51 ms` faster than `89.09%` submissions  
**Memory usage:** `51.4 MB` less than `89.81%` submissions  

``` typescript
function generateParenthesis(n: number): string[] {
    let result: string[] = [];
    recGenerator("", 0, 0, n, result);
    return result;
}

function recGenerator(currentString: string, opened: number, closed: number, n: number, result: string[]): void {
    if (currentString.length == 2*n) {
        result.push(currentString);
        return;
    }
    if (opened < n) {
        recGenerator(currentString + "(", opened + 1, closed, n, result);
    }
    if (closed < opened) {
        recGenerator(currentString + ")", opened, closed + 1, n, result);
    }
}
```

### 3. GO

**Runtime:** `2 ms` faster than `75.14%` submissions  
**Memory usage:** `2 MB` less than `82.58%` submissions  

``` go
func generateParenthesis(n int) []string {
	result := []string{}
	recGenerator("", 0, 0, n, &result)
	return result
}

func recGenerator(currentString string, opened int, closed int, n int, result *[]string) {
	if len(currentString) == 2*n {
		*result = append(*result, currentString)
		return
	}
	if opened < n {
		recGenerator(currentString+"(", opened+1, closed, n, result)
	}
	if closed < opened {
		recGenerator(currentString+")", opened, closed+1, n, result)
	}
}
```