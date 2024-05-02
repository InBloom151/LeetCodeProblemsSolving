function myAtoi(s: string): number {
    let sign = 1, result = 0, i = 0;
    const INT_MAX = 2 ** 31 - 1, INT_MIN = -(2 ** 31);

    while (s[i] === ' ') i++;
    if (s[i] === '-') { sign = -1; i++; }
    else if (s[i] === '+') i++;

    for (; i < s.length && '0' <= s[i] && s[i] <= '9'; i++) {
        result = result * 10 + (s.charCodeAt(i) - 48);
        if (sign * result > INT_MAX) return INT_MAX;
        if (sign * result < INT_MIN) return INT_MIN;
    }

    return sign * result;
}

console.log(myAtoi("42"))
console.log(myAtoi(" -042"))
console.log(myAtoi("1337c0d3"))
console.log(myAtoi("0-1"))
console.log(myAtoi("words and 987"))