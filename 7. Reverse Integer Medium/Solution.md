### This approach has a time complexity of `O(log(n))`


### 1. Python

**Runtime:** `31 ms` faster than `86.50%` submissions  
**Memory usage:** `16.6 MB` less than `68.88%` submissions  

``` python
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
```

### 2. TypeScript

**Runtime:** `67 ms` faster than `92.53%` submissions  
**Memory usage:** `53.8 MB` less than `36.83%` submissions  

``` typescript
function reverse(x: number): number {
    let isNegative = x < 0;
    let reversedNumber = 0;

    if (isNegative) {
        x = -x;
    }

    while (x) {
        let digit = x % 10;
        if (reversedNumber > (2 ** 31 - 1) / 10 || (reversedNumber === (2 ** 31 - 1) / 10 && digit > 7)) {
            return 0;
        }
        reversedNumber = reversedNumber * 10 + digit;
        x = Math.floor(x / 10);
    }

    return isNegative ? -reversedNumber : reversedNumber;
};
```

### 3. GO

**Runtime:** `0 ms` faster than `100.00%` submissions  
**Memory usage:** `2.3 MB` less than `79.18%` submissions  

``` go
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
	isNegative := x < 0
	x = int(math.Abs(float64(x)))
	reversedNumber := 0

	for x > 0 {
		digit := x % 10
		reversedNumber = reversedNumber*10 + digit
		x /= 10
	}

	if reversedNumber > math.MaxInt32 {
		return 0
	}

	if isNegative {
		return -reversedNumber
	}
	return reversedNumber
}
```

### 4. Java

**Runtime:** `1 ms` faster than `82.89%` submissions  
**Memory usage:** `41.1 MB` less than `22.73%` submissions  

``` java
public class solution {
    public static int reverse(int x) {
        int reversedNumber = 0;

        while (x != 0) {
            int digit = x % 10;
            if (reversedNumber > Integer.MAX_VALUE / 10 || (reversedNumber == Integer.MAX_VALUE / 10 && digit > 7)) {
                return 0;
            }
            if (reversedNumber < Integer.MIN_VALUE / 10 || (reversedNumber == Integer.MIN_VALUE / 10 && digit < -8)) {
                return 0;
            }
            reversedNumber = reversedNumber * 10 + digit;
            x /= 10;
        }

        return reversedNumber;
    }
}
```