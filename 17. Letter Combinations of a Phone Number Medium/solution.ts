function letterCombinations(digits: string): string[] {
    const PHONE_NUMS: { [key: string]: string } = {
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    };

    if (digits === "") {
        return [];
    }

    const result: string[] = [];
    getCombinations("", 0, digits, result, PHONE_NUMS);
    return result;
}

function getCombinations(
    combination: string,
    index: number,
    digits: string,
    result: string[],
    associations: { [key: string]: string }
): void {
    if (index === digits.length) {
        result.push(combination);
        return;
    }

    const currentDigit: string = digits[index];
    for (const letter of associations[currentDigit]) {
        getCombinations(combination + letter, index + 1, digits, result, associations);
    }
}

console.log(letterCombinations("23"));
console.log(letterCombinations(""));
console.log(letterCombinations("2"));