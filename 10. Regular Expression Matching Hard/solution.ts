function isMatch(s: string, p: string): boolean {
    const memo: {[key: string]: boolean} = {};

    function dp(i: number, j: number): boolean {
        const key = `${i},${j}`;
        if (memo[key] !== undefined) {
            return memo[key];
        }

        if (j === p.length) {
            const ans = i === s.length;
            memo[key] = ans;
            return ans;
        } else {
            const firstMatch = i < s.length && (p[j] === s[i] || p[j] === '.');

            if (j + 1 < p.length && p[j + 1] === '*') {
                const ans = dp(i, j + 2) || (firstMatch && dp(i + 1, j));
                memo[key] = ans;
                return ans;
            } else {
                const ans = firstMatch && dp(i + 1, j + 1);
                memo[key] = ans;
                return ans;
            }
        }
    }

    return dp(0, 0);
}

console.log(isMatch('aa', 'a'));
console.log(isMatch('aa', 'a*'));
console.log(isMatch('ab', '.*'));