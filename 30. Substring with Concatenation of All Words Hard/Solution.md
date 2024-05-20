### This approach has a time complexity of `O(log(n))`

### 1. Python

**Runtime:** `62 ms` faster than `97.15%` submissions  
**Memory usage:** `17.4 MB` less than `40.83` submissions  

``` python
from collections import defaultdict

def find_substring(s: str, words: list[str]) -> list[int]:
    if not words[0]:
        return []
    
    word_length = len(words[0])
    word_count = len(words)
    word_frequency = defaultdict(int)
    
    for word in words:
        word_frequency[word] += 1
    
    result = []
    
    for i in range(word_length):
        left = i
        right = i
        current_count = 0
        current_frequency = {}
        
        while right + word_length <= len(s):
            word = s[right:right + word_length]
            right += word_length
            
            if word in word_frequency:
                if word in current_frequency:
                    current_frequency[word] += 1
                else:
                    current_frequency[word] = 1
                
                current_count += 1
                
                while current_frequency[word] > word_frequency[word]:
                    left_word = s[left:left + word_length]
                    current_frequency[left_word] -= 1
                    left += word_length
                    current_count -= 1
                
                if current_count == word_count:
                    result.append(left)
            else:
                current_frequency.clear()
                current_count = 0
                left = right
    
    return result
```

### 2. TypeScript

**Runtime:** `90 ms` faster than `89.49%` submissions  
**Memory usage:** `59.6 MB` less than `26.75%` submissions  

``` typescript
function findSubstring(s: string, words: string[]): number[] {
    if (!words[0]) {
        return [];
    }

    const wordLength = words[0].length;
    const wordCount = words.length;
    let wordFrequency: {[key: string]: number} = {};

    for (let word of words) {
        if (wordFrequency[word]) {
            wordFrequency[word]++;
        } else {
            wordFrequency[word] = 1;
        }
    }

    let result: number[] = [];

    for (let i = 0; i < wordLength; i++) {
        let left = i;
        let right = i;
        let currentCount = 0;
        let currentFrequency: {[key: string]: number} = {};

        while (right + wordLength <= s.length) {
            let word = s.slice(right, right + wordLength);
            right += wordLength;

            if (wordFrequency[word]) {
                if (currentFrequency[word]) {
                    currentFrequency[word]++;
                } else {
                    currentFrequency[word] = 1;
                }

                currentCount++;

                while (currentFrequency[word] > wordFrequency[word]) {
                    let leftWord = s.slice(left, left + wordLength);
                    currentFrequency[leftWord]--;
                    left += wordLength;
                    currentCount--;
                }

                if (currentCount === wordCount) {
                    result.push(left);
                }
            } else {
                currentFrequency = {};
                currentCount = 0;
                left = right;
            }
        }

    }
    return result;
}
```

### 3. GO

**Runtime:** `7 ms` faster than `96.53%` submissions  
**Memory usage:** `6.5 MB` less than `57.07%` submissions  

``` go
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	wordLength := len(words[0])
	wordCount := len(words)
	wordFrequency := map[string]int{}

	for _, word := range words {
		wordFrequency[word]++
	}

	result := []int{}

	for i := 0; i < wordLength; i++ {
		left := i
		right := i
		currentCount := 0
		currentFrequency := map[string]int{}

		for right+wordLength <= len(s) {
			word := s[right : right+wordLength]
			right += wordLength

			if wordFrequency[word] != 0 {
				currentFrequency[word]++
				currentCount++

				for currentFrequency[word] > wordFrequency[word] {
					leftWord := s[left : left+wordLength]
					currentFrequency[leftWord]--
					left += wordLength
					currentCount--
				}

				if currentCount == wordCount {
					result = append(result, left)
				}
			} else {
				currentFrequency = map[string]int{}
				currentCount = 0
				left = right
			}
		}
	}

	return result
}
```