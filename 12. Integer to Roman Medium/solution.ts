function intToRoman(num: number): string {
    const romanNumerals: [number, string][] = [
        [1000, "M"],
        [900, "CM"],
        [500, "D"],
        [400, "CD"],
        [100, "C"],
        [90, "XC"],
        [50, "L"],
        [40, "XL"],
        [10, "X"],
        [9, "IX"],
        [5, "V"],
        [4, "IV"],
        [1, "I"],
    ];

    let result: string = '';

    for (let [value, roman] of romanNumerals) {
        while (num >= value) {
            result += roman;
            num -= value;
        }
        if (num === 0) break;
    }

    return result;
}   

console.log(intToRoman(3749))
console.log(intToRoman(58))
console.log(intToRoman(1994))