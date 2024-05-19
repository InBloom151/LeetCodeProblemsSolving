### This approach has a time complexity of `O(log(n))`

### 1. Python

**Runtime:** `30 ms` faster than `89.87%` submissions  
**Memory usage:** `16.5 MB` less than `75.29` submissions  

``` python
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
```

### 2. TypeScript

**Runtime:** `66 ms` faster than `95.05%` submissions  
**Memory usage:** `53.5 MB` less than `74.73%` submissions  

``` typescript
function divide(dividend: number, divisor: number): number {
    const MAX = 2147483647;
    const MIN = -2147483648;
    if (divisor === 0 || (dividend === MIN && divisor === -1)) {
        return MAX;
    }
    if (divisor === dividend) {
        return 1;
    }
    const sign = (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) ? -1 : 1;
    let quotient = 0;
    let absoluteDividend = Math.abs(dividend);
    let absoluteDivisor = Math.abs(divisor);
    while (absoluteDividend >= absoluteDivisor) {
        let shift = 0;
        let shiftedDivisor = absoluteDivisor;
        while (absoluteDividend >= shiftedDivisor) {
            shift++;
            shiftedDivisor = absoluteDivisor << shift;
            if (shiftedDivisor < 0) {
                break;
            }
        }
        quotient += (1 << (shift - 1));
        absoluteDividend -= absoluteDivisor << (shift - 1);
    }
    return sign === -1 ? -quotient : quotient;
};
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.5 MB` less than `97.76%` submissions  

``` go
func divide(dividend int, divisor int) int {
    MIN := -2147483648
    MAX := 2147483647

    if divisor == 0 || (dividend == MIN && divisor == -1) {
        return MAX
    }

    if dividend == 0 {
        return 0
    }

    if divisor == 1 {
        return dividend
    }

    if divisor == -1 {
        if dividend == MIN {
            return MAX
        }
        return -dividend
    }

    var sign int
    if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        sign = -1
    } else {
        sign = 1
    }

    absoluteDividend := int(math.Abs(float64(dividend)))
    absoluteDivisor := int(math.Abs(float64(divisor)))

    quotient := 0

    for absoluteDividend >= absoluteDivisor {
        shift := 0
        tempDivisor := absoluteDivisor

        for absoluteDividend >= tempDivisor<<1 {
            shift++
            tempDivisor <<= 1
        }

        quotient += 1 << shift
        absoluteDividend -= tempDivisor
    }

    if sign == -1 {
        return -quotient
    }
    return quotient
}
```