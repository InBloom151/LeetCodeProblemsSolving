def reverse(x: int) -> int:
    is_negative = x < 0
    x = abs(x)
    reversed_number = 0
    
    while x:
        digit = x % 10
        reversed_number = reversed_number * 10 + digit
        x //= 10
    
    if reversed_number > 2 ** 31 - 1:
        return 0
    
    return -reversed_number if is_negative else reversed_number

print(reverse(123))
print(reverse(-123))
print(reverse(120))