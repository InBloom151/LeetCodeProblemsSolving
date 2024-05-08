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

print(is_valid("()"))
print(is_valid("()[]{}"))
print(is_valid("(]"))