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

console.log(reverse(123))
console.log(reverse(-123))
console.log(reverse(120))