function romanToInt(s: string): number {
    const ASSOCIATIONS: {[key: string]: number} = {
        'I': 1, 
        'V': 5, 
        'X': 10, 
        'L': 50,
        'C': 100, 
        'D': 500, 
        'M': 1000,
    }

    let total = 0;
    let prevValue = 0;

    for (let numeral = s.length - 1; numeral >= 0; numeral--) {
        let value = ASSOCIATIONS[s[numeral]];

        total = value < prevValue ? total -= value : total += value;

        prevValue = value;
    }

    return total;
}

console.log(romanToInt("III"))
console.log(romanToInt("LVIII"))
console.log(romanToInt("MCMXCIV"))