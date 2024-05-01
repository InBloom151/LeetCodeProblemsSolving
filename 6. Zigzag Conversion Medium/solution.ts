function convert(s: string, numRows: number): string {
    if (numRows === 1 || numRows >= s.length) {
        return s;
    }

    const result = [];
    const stepSize = 2 * numRows - 2;
    const len = s.length;

    for (let i = 0; i < numRows; i++) {
        let index = i;
        let step1 = 2 * (numRows - 1 - i);
        let step2 = stepSize - step1;

        while (index < len) {
            if (step1 !== 0) {
                result.push(s[index]);
                index += step1;
            }
            if (step2 !== 0 && index < len) {
                result.push(s[index]);
                index += step2;
            }
        }
    }

    return result.join("");
}

console.log(convert("PAYPALISHIRING", 3))
console.log(convert("PAYPALISHIRING", 4))
console.log(convert("A", 1))