def divide(dividend: int, divisor: int) -> int:
    INT_MAX = 2**31 - 1
    INT_MIN = -2**31

    if dividend == INT_MIN and divisor == -1:
        return INT_MAX
    
    negative = (dividend < 0) ^ (divisor < 0)

    dividend = abs(dividend)
    divisor = abs(divisor)

    quotient = 0

    while dividend >= divisor:
        temp_divisor, multiple = divisor, 1
        while dividend >= (temp_divisor << 1):
            temp_divisor <<= 1
            multiple <<= 1
        dividend -= temp_divisor
        quotient += multiple

    if negative:
        quotient = -quotient

    return min(max(INT_MIN, quotient), INT_MAX)

print(divide(10, 3))
print(divide(7, -3))