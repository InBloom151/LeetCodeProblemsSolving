function generateParenthesis(n: number): string[] {
    let result: string[] = [];
    recGenerator("", 0, 0, n, result);
    return result;
}

function recGenerator(currentString: string, opened: number, closed: number, n: number, result: string[]): void {
    if (currentString.length == 2*n) {
        result.push(currentString);
        return;
    }
    if (opened < n) {
        recGenerator(currentString + "(", opened + 1, closed, n, result);
    }
    if (closed < opened) {
        recGenerator(currentString + ")", opened, closed + 1, n, result);
    }
}

console.log(generateParenthesis(3))
console.log(generateParenthesis(1))