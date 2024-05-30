### This approach has a time complexity of `O(2^n)`

### 1. Python

**Runtime:** `44 ms` faster than `90.62%` submissions  
**Memory usage:** `16.5 MB` less than `97.98%` submissions  

``` python
def combination_sum(candidates: list[int], target: int) -> list[list[int]]:
    result = []

    def backtrack(remain: int, combo: list[int], start: int):
        if remain == 0:
            result.append(list(combo))
            return
        elif remain < 0:
            return
        
        for i in range(start, len(candidates)):
            combo.append(candidates[i])
            backtrack(remain - candidates[i], combo, i)
            combo.pop()

    backtrack(target, [], 0)
    return result
```

### 2. TypeScript

**Runtime:** `71 ms` faster than `90.69%` submissions  
**Memory usage:** `54.5 MB` less than `77.75%` submissions  

``` typescript
function combinationSum(candidates: number[], target: number): number[][] {
    let result: number[][] = [];

    function backtrack(remain: number, combo: number[], start: number): void {
        if (remain === 0) {
            result.push(Array.from(combo));
            return;
        } else if (remain < 0) {
            return;
        }

        for (let i = start; i < candidates.length; i++) {
            combo.push(candidates[i]);
            backtrack(remain - candidates[i], combo, i);
            combo.pop();
        }
    }

    backtrack(target, [], 0);
    return result;
}
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `3.2 MB` less than `78.70%` submissions  

``` go
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	var backtrack func(remain int, combo []int, start int)
	backtrack = func(remain int, combo []int, start int) {
		if remain == 0 {
			combination := make([]int, len(combo))
			copy(combination, combo)
			result = append(result, combination)
			return
		} else if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			combo = append(combo, candidates[i])
			backtrack(remain-candidates[i], combo, i)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(target, []int{}, 0)
	return result
}
```