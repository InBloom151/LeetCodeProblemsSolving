def myAtoi(s: str) -> int:
    s = s.strip()
    
    if not s:
        return 0
    
    sign = 1
    if s[0] == '-':
        sign = -1
        s = s[1:]
    elif s[0] == '+':
        s = s[1:]
    
    result = 0
    for char in s:
        if not char.isdigit():
            break
        result = result * 10 + int(char)
    
    result *= sign
    
    INT_MAX = 2**31 - 1
    INT_MIN = -2**31
    if result > INT_MAX:
        return INT_MAX
    elif result < INT_MIN:
        return INT_MIN
    
    return result

print(myAtoi("42"))
print(myAtoi(" -042"))
print(myAtoi("1337c0d3"))
print(myAtoi("0-1"))
print(myAtoi("words and 987"))
