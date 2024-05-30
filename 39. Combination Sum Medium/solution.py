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

candidates1 = [2, 3, 6, 7]
target1 = 7
print(combination_sum(candidates1, target1))
candidates2 = [2, 3, 5]
target2 = 8
print(combination_sum(candidates2, target2))
candidates3 = [2]
target3 = 1
print(combination_sum(candidates3, target3))