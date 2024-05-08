function isValid(s: string): boolean {
    const result: string[] = [];
    const openBrackets = ['(', '[', '{'];
    const closingBrackets = [')', ']', '}'];

    for (let bracket of s) {
        if (openBrackets.includes(bracket)) {
            result.push(bracket);
        } else if (closingBrackets.includes(bracket)) {
            const expectedBracket = openBrackets[closingBrackets.indexOf(bracket)];
            if (result.pop() !== expectedBracket) return false;
        }
    }

    return result.length === 0;
}

console.log(isValid("()"))
console.log(isValid("()[]{}"))
console.log(isValid("(]"))