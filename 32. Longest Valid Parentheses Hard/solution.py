def longest_valid_parentheses(s: str) -> int:
    stack = [-1]
    result = 0
    for i, c in enumerate(s):
        if c == "(":
            stack.append(i)
        else:
            stack.pop()
            if not stack:
                stack.append(i)
            else:
                result = max(result, i - stack[-1])
    return result

print(longest_valid_parentheses("(()"))
print(longest_valid_parentheses(")()())"))
print(longest_valid_parentheses(""))
print(longest_valid_parentheses("((()))"))
print(longest_valid_parentheses("((()))()"))
print(longest_valid_parentheses("(()()())()((()))"))