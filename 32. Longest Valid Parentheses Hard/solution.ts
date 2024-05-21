function longestValidParentheses(s: string): number {
    let stack: number[] = [-1];
    let result = 0;

    for (let i = 0; i < s.length; i++) {
        if (s.charAt(i) == "(") {
            stack.push(i);
        } else {
            stack.pop();
            if (stack.length == 0) {
                stack.push(i);
            } else {
                result = Math.max(result, i - stack[stack.length-1]);
            }
        }
    }

    return result;
}

console.log(longestValidParentheses("(()"));
console.log(longestValidParentheses(")()())"));
console.log(longestValidParentheses(""));
console.log(longestValidParentheses("((()))"));
console.log(longestValidParentheses("((()))()"));
console.log(longestValidParentheses("(()()())()((()))"));