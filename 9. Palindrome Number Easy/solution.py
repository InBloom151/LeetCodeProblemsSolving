def is_palindrome(x: int) -> bool:
    return str(x) == ''.join(reversed(str(x)))


print(is_palindrome(121))
print(is_palindrome(-121))
print(is_palindrome(10))