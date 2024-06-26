def is_match(s: str, p: str) -> bool:
    memo = {}
    
    def dp(i: int, j: int) -> bool:
        if (i, j) in memo:
            return memo[(i, j)]

        if j == len(p):
            ans = i == len(s)
        else:
            first_match = i < len(s) and p[j] in {s[i], '.'}

            if j + 1 < len(p) and p[j + 1] == '*':
                ans = dp(i, j + 2) or (first_match and dp(i + 1, j))
            else:
                ans = first_match and dp(i + 1, j + 1)

        memo[(i, j)] = ans
        return ans

    return dp(0, 0)
    

print(is_match('aa', 'a'))
print(is_match('aa', 'a*'))
print(is_match('ab', '.*'))