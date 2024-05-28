### This approach has a time complexity of `O(2^n)`

### 1. Python

**Runtime:** `34 ms` faster than `95.77%` submissions  
**Memory usage:** `16.6 MB` less than `81.10%` submissions  

``` python
def count_and_say(n: int) -> str:
    if n == 1:
        return "1"
    
    def next_sequence(sequence: str) -> str:
        result = []
        count = 1
        char = sequence[0]

        for i in range(1, len(sequence)):
            if sequence[i] == char:
                count += 1
            else:
                result.append(str(count))
                result.append(char)
                char = sequence[i]
                count = 1
        result.append(str(count))
        result.append(char)

        return "".join(result)
    
    current_sequence = "1"
    for _ in range(1, n):
        current_sequence = next_sequence(current_sequence)

    return current_sequence
```

### 2. TypeScript

**Runtime:** `55 ms` faster than `85.71%` submissions  
**Memory usage:** `52.6 MB` less than `40.87%` submissions  

``` typescript
function countAndSay(n: number): string {
    if (n === 1) return "1";

    function nextSequence(sequence: string): string {
        let result: string[] = [];
        let count = 1;
        let char = sequence.charAt(0);

        for (let i = 1; i < sequence.length; i++) {
            if (sequence.charAt(i) === char) {
                count++;
            } else {
                result.push(count.toString());
                result.push(char);
                char = sequence.charAt(i);
                count = 1;
            }
        }
        result.push(count.toString());
        result.push(char);

        return result.join("");
    }

    let currentSequence = "1";
    for (let i = 1; i < n; i++) {
        currentSequence = nextSequence(currentSequence);
    }
    return currentSequence;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.7 MB` less than `80.99%` submissions  

``` go
import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	var nextSequence func(sequence string) string
	nextSequence = func(sequence string) string {
		result := []byte{}
		count := 1
		char := sequence[0]

		for i := 1; i < len(sequence); i++ {
			if sequence[i] == char {
				count++
			} else {
				result = append(result, []byte(strconv.Itoa(count))...)
				result = append(result, char)
				char = sequence[i]
				count = 1
			}
		}
		result = append(result, []byte(strconv.Itoa(count))...)
		result = append(result, char)

		return string(result)
	}

	currentSequence := "1"
	for i := 1; i < n; i++ {
		currentSequence = nextSequence(currentSequence)
	}

	return currentSequence
}
```